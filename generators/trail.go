// +build !js

package generators

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/influx6/faux/fmtwriter"

	"github.com/influx6/moz/ast"
	"github.com/influx6/moz/gen"
	"github.com/influx6/trail/generators/data"
)

var (
	inGOPATH    = os.Getenv("GOPATH")
	inGOPATHSrc = filepath.Join(inGOPATH, "src")
	badSymbols  = regexp.MustCompile(`[(|\-|_|\W|\d)+]`)
	notAllowed  = regexp.MustCompile(`[^(_|\w|\d)+]`)
)

// TrailPackages returns a slice of WriteDirectives which contain data to be written to disk to create
// a suitable package for asset bundle.
func TrailPackages(an ast.AnnotationDeclaration, pkg ast.PackageDeclaration, pkgDir ast.Package) ([]gen.WriteDirective, error) {
	if len(an.Arguments) == 0 {
		return nil, errors.New("Expected atleast one argument for annotation as component name")
	}

	workDir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve current directory path: %+q", err)
	}

	gridCSSData := data.Must("grid.css.gen")
	gridNormCSS := data.Must("normalize.css.gen")

	packageDir, err := filepath.Rel(inGOPATHSrc, workDir)
	if err != nil {
		fmt.Printf("Failed to retrieve package directory path in go src: %+q\n", err)
	}

	componentName := badSymbols.ReplaceAllString(an.Arguments[0], "")
	componentNameLower := strings.ToLower(componentName)
	componentPackageDir := filepath.Join(packageDir, componentNameLower)

	var targetPkg string

	if componentNameLower != "" {
		targetPkg = componentNameLower
	} else {
		targetPkg = filepath.Base(workDir)
	}

	publicStandInGen := gen.Block(
		gen.Package(
			gen.Name(targetPkg),
			gen.SourceText(
				string(data.Must("bundle.gen")),
				struct {
					Name    string
					Package string
				}{
					Name:    componentName,
					Package: componentNameLower,
				},
			),
		),
	)

	publicGen := gen.Block(
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
				Settings:      true,
				Name:          componentName,
				Package:       componentNameLower,
				LessFile:      fmt.Sprintf("less/%s.less", componentNameLower),
			},
		),
	)

	lessGen := gen.Block(
		gen.SourceText(
			string(data.Must("main.less.gen")),
			struct {
				Name    string
				Package string
			}{
				Name:    componentName,
				Package: componentNameLower,
			},
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

	tomlGen := gen.Block(
		gen.SourceText(
			string(data.Must("settings.toml.gen")),
			struct {
				Name    string
				Package string
			}{
				Name:    targetPkg,
				Package: componentPackageDir,
			},
		),
	)

	lessName := "index"
	if componentName != "" {
		lessName = componentNameLower
	}

	commands := []gen.WriteDirective{
		{
			DontOverride: false,
			Dir:          componentNameLower,
		},
		{
			DontOverride: false,
			Writer:       htmlGen,
			FileName:     "index.html",
			Dir:          componentNameLower,
		},
		{
			DontOverride: false,
			FileName:     "main.js",
			Dir:          filepath.Join(componentNameLower, "js"),
			Writer:       bytes.NewBufferString("//strictmode"),
		},
		{
			DontOverride: false,
			Dir:          filepath.Join(componentNameLower, "css"),
			FileName:     "normalize.css",
			Writer:       bytes.NewBuffer(gridNormCSS),
		},
		{
			DontOverride: false,
			Dir:          filepath.Join(componentNameLower, "css"),
			FileName:     "grid.css",
			Writer:       bytes.NewBuffer(gridCSSData),
		},
		{
			DontOverride: false,
			Dir:          filepath.Join(componentNameLower, "less"),
		},
		{
			DontOverride: false,
			Writer:       lessGen,
			Dir:          filepath.Join(componentNameLower, "less"),
			FileName:     fmt.Sprintf("%s.less", lessName),
		},
		{
			DontOverride: true,
			Writer:       tomlGen,
			Dir:          componentNameLower,
			FileName:     "settings.toml",
		},
		{
			DontOverride: true,
			Dir:          componentNameLower,
			FileName:     "generate.go",
			Writer:       fmtwriter.New(publicGen, true, true),
		},
		{
			DontOverride: true,
			Dir:          componentNameLower,
			FileName:     "bundle.go",
			Writer:       fmtwriter.New(publicStandInGen, true, true),
		},
	}

	return commands, nil
}

func validateName(val string) bool {
	return notAllowed.MatchString(val)
}
