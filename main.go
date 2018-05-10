package main

import (
	"flag"
	"os"

	"github.com/gokit/assetkit/generators"
	"github.com/influx6/faux/flags"
	"github.com/influx6/moz/ast"
)

func main() {
	flags.Run("assetkit", flags.Command{
		Name:      "public",
		ShortDesc: "Generates bundling for public files",
		Desc:      "Generates asset bundling for standard public static files",
		Action: func(ctx flags.Context) error {
			force := ctx.GetBool("force")

			name := flag.Arg(1)

			currentdir, err := os.Getwd()
			if err != nil {
				return err
			}

			commands, err := generators.TrailPackages(
				ast.AnnotationDeclaration{Arguments: []string{name}},
				ast.PackageDeclaration{FilePath: currentdir},
				ast.Package{},
			)
			if err != nil {
				return err
			}

			return ast.SimpleWriteDirectives("", force, commands...)
		},
		Flags: []flags.Flag{
			&flags.BoolFlag{
				Name: "force",
				Desc: "force regeneration of packages annotation directives.",
			},
		},
	},
		flags.Command{
			Name:      "view",
			ShortDesc: "Generates bundling with a html file",
			Desc:      "Generates asset bundling isolated view package",
			Action: func(ctx flags.Context) error {
				force := ctx.GetBool("force")

				name := flag.Arg(1)

				currentdir, err := os.Getwd()
				if err != nil {
					return err
				}

				commands, err := generators.TrailView(
					ast.AnnotationDeclaration{Arguments: []string{name}},
					ast.PackageDeclaration{FilePath: currentdir},
					ast.Package{},
				)
				if err != nil {
					return err
				}

				return ast.SimpleWriteDirectives("", force, commands...)
			},
			Flags: []flags.Flag{
				&flags.BoolFlag{
					Name: "force",
					Desc: "force regeneration of packages annotation directives.",
				},
			},
		},
		flags.Command{
			Name:      "static",
			Desc:      "Generates bundling for general static files",
			ShortDesc: "Generates bundling general use case static files",
			Action: func(ctx flags.Context) error {
				force := ctx.GetBool("force")
				name := flag.Arg(1)

				currentdir, err := os.Getwd()
				if err != nil {
					return err
				}

				commands, err := generators.TrailFiles(
					ast.AnnotationDeclaration{Arguments: []string{name}},
					ast.PackageDeclaration{FilePath: currentdir},
					ast.Package{},
				)
				if err != nil {
					return err
				}

				return ast.SimpleWriteDirectives("", force, commands...)
			},
			Flags: []flags.Flag{
				&flags.BoolFlag{
					Name: "force",
					Desc: "force regeneration of packages annotation directives.",
				},
			},
		})
}
