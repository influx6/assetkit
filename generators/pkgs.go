package generators

import (
	"errors"
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

	componentName := badSymbols.ReplaceAllString(an.Arguments[0], "")
	componentNameLower := strings.ToLower(componentName)

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
				Name:          componentName,
				Package:       componentNameLower,
				TargetPackage: componentNameLower,
			},
		),
	)

	pipeGen := gen.Block(
		gen.Package(
			gen.Name(componentName),
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

	return []gen.WriteDirective{
		{
			DontOverride: true,
			FileName:     "bundle.go",
			Dir:          componentNameLower,
			Writer:       fmtwriter.New(pipeGen, true, true),
		},
		{
			DontOverride: true,
			FileName:     "generate.go",
			Dir:          componentNameLower,
			Writer:       fmtwriter.New(generatorGen, true, true),
		},
	}, nil
}
