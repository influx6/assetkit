package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/influx6/moz/ast"
	"github.com/influx6/trail/generators"
)

var (
	version   = "0.0.1" // rely on linker -ldflags -X main.version"
	gitCommit = ""      // rely on linker: -ldflags -X main.gitCommit"
)

var (
	getVersion = *flag.Bool("v", false, "Print version")
)

func main() {
	flag.Usage = printUsage
	flag.Parse()

	// if we are to print getVersion.
	if getVersion {
		printVersion()
		return
	}

	name := flag.Arg(0)
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get directory path: %+q", err)
		return
	}

	commands, err := generators.TrailPackages(
		ast.AnnotationDeclaration{Arguments: []string{name}},
		ast.PackageDeclaration{FilePath: currentDir},
		ast.Package{},
	)
	if err != nil {
		log.Fatalf("Failed to generate trail directives: %+q", err)
		return
	}

	if err := ast.SimpleWriteDirectives(name, true, commands...); err != nil {
		log.Fatalf("Failed to create package directories: %+q", err)
		return
	}

	log.Println("Trail asset bundling ready!")
}

// printVersion prints corresponding build getVersion with associated build stamp and git commit if provided.
func printVersion() {
	fragments := []string{version}

	if gitCommit != "" {
		fragments = append(fragments, fmt.Sprintf("git#%s", gitCommit))
	}

	fmt.Fprint(os.Stdout, strings.Join(fragments, " "))
}

// printUsage prints out usage message for CLI tool.
func printUsage() {
	fmt.Fprintf(os.Stdout, `Usage: trail [options]
Trail creates a package for package of web assets using it's internal bundlers.

EXAMPLES:

	trail static-data

FLAGS:

  -v          Print version.

`)
}
