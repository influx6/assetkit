package generators

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/influx6/faux/fmtwriter"

	"github.com/influx6/moz/ast"
	"github.com/influx6/moz/gen"
	"github.com/influx6/trail/generators/data"
)

// TrailView returns a series of file commands which create asset bundling for a giving file.
func TrailView(an ast.AnnotationDeclaration, pkg ast.PackageDeclaration, pk ast.Package) ([]gen.WriteDirective, error) {
	if len(an.Arguments) == 0 {
		return nil, errors.New("Expected atleast one argument for annotation as component name")
	}

	workDir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve current directory path: %+q", err)
	}

	componentName := badSymbols.ReplaceAllString(an.Arguments[0], "")
	componentNameLower := strings.ToLower(componentName)

	var targetPkg string

	if componentName != "" {
		targetPkg = strings.ToLower(componentName)
	} else {
		targetPkg = filepath.Base(workDir)
	}

	generatorGen := gen.Block(
		gen.SourceText(
			string(data.Must("pack-bundle.gen")),
			struct {
				Name          string
				LessFile      string
				Package       string
				TargetDir     string
				TargetPackage string
				Settings      bool
			}{
				TargetDir:     "./",
				TargetPackage: targetPkg,
				Name:          componentName,
				Package:       componentNameLower,
			},
		),
	)

	pipeGen := gen.Block(
		gen.Package(
			gen.Name(targetPkg),
			gen.Block(
				gen.Text("\n"),
				gen.Text("//go:generate go run generate.go"),
				gen.Text("\n"),
				gen.SourceText(
					string(data.Must("bundle.gen")),
					nil,
				),
			),
		),
	)

	htmlGen := gen.Block(
		gen.SourceText(
			string(data.Must("base.html.gen")),
			struct {
				Name   string
				Path   string
				JSFile string
			}{
				Name:   targetPkg,
				Path:   "public",
				JSFile: fmt.Sprintf("%s/%s", "js", "main.js"),
			},
		),
	)

	return []gen.WriteDirective{
		{
			DontOverride: false,
			FileName:     "bundle.go",
			Dir:          componentNameLower,
			Writer:       fmtwriter.New(pipeGen, true, true),
		},
		{
			DontOverride: false,
			Writer:       htmlGen,
			Dir:          componentNameLower,
			FileName:     fmt.Sprintf("%s.html", targetPkg),
		},
		{
			DontOverride: false,
			FileName:     "generate.go",
			Dir:          componentNameLower,
			Writer:       fmtwriter.New(generatorGen, true, true),
		},
	}, nil
}
