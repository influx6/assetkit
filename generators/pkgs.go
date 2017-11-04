package generators

// import (
// 	"errors"
// 	"fmt"
// 	"strings"
// 	"text/template"
//
// 	"github.com/influx6/faux/fmtwriter"
//
// 	"github.com/gu-io/gu/generators/data"
// 	"github.com/influx6/moz/ast"
// 	"github.com/influx6/moz/gen"
// )
//
// // ComponentPackageGenerator which defines a  function for generating a type for receiving a giving
// //	struct type has a notification type which can then be wired as a notification.EventDistributor.
// //
// //	Annotation: @notification:event
// func ComponentPackageGenerator(an ast.AnnotationDeclaration, pkg ast.PackageDeclaration, pk ast.Package) ([]gen.WriteDirective, error) {
// 	if len(an.Arguments) == 0 {
// 		return nil, errors.New("Expected atleast one argument for annotation as component name")
// 	}
//
// 	componentName := an.Arguments[0]
// 	componentNameLower := strings.ToLower(componentName)
//
// 	typeGen := gen.Block(
// 		gen.Package(
// 			gen.Name(componentName),
// 			gen.Imports(
// 				gen.Import("github.com/gu-io/gu", ""),
// 				gen.Import("github.com/gu-io/gu/trees", ""),
// 				gen.Import("github.com/gu-io/gu/trees/elems", ""),
// 				gen.Import("github.com/gu-io/gu/trees/property", ""),
// 			),
// 			gen.Block(
// 				gen.SourceTextWith(
// 					string(data.Must("scaffolds/component.gen")),
// 					template.FuncMap{},
// 					struct {
// 						Name string
// 					}{
// 						Name: componentName,
// 					},
// 				),
// 			),
// 		),
// 	)
//
// 	generatorGen := gen.Block(
// 		gen.SourceText(
// 			string(data.Must("scaffolds/pack-bundle.gen")),
// 			struct {
// 				Name          string
// 				LessFile      string
// 				Package       string
// 				TargetDir     string
// 				TargetPackage string
// 			}{
// 				TargetDir:     "./",
// 				Name:          componentName,
// 				Package:       componentNameLower,
// 				TargetPackage: componentNameLower,
// 			},
// 		),
// 	)
//
// 	pipeGen := gen.Block(
// 		gen.Package(
// 			gen.Name(componentName),
// 			gen.Block(
// 				gen.Text("\n"),
// 				gen.Text("//go:generate go run generate.go"),
// 				gen.Text("\n"),
// 				gen.SourceText(
// 					string(data.Must("scaffolds/bundle.gen")),
// 					nil,
// 				),
// 			),
// 		),
// 	)
//
// 	return []gen.WriteDirective{
// 		{
// 			DontOverride: true,
// 			Dir:          componentNameLower,
// 			FileName:     fmt.Sprintf("%s.go", componentNameLower),
// 			Writer:       fmtwriter.New(typeGen, true, true),
// 		},
// 		{
// 			DontOverride: true,
// 			FileName:     fmt.Sprintf("%s_bundle.go", componentNameLower),
// 			Dir:          componentNameLower,
// 			Writer:       fmtwriter.New(pipeGen, true, true),
// 		},
// 		{
// 			DontOverride: true,
// 			FileName:     "generate.go",
// 			Dir:          componentNameLower,
// 			Writer:       fmtwriter.New(generatorGen, true, true),
// 		},
// 	}, nil
// }
