AssetKit
---------
[![Go Report Card](https://goreportcard.com/badge/github.com/gokit/assetkit)](https://goreportcard.com/report/github.com/gokit/assetkit)

AssetKit code generates asset bundling packages for you. 


## Install

```
go get github.com/gokit/assetkit
```

## Usage

Simple call the `assetkit` CLI with a wanted asset package:

```bash
> assetkit
Usage: assetkit [flags] [command] 

⡿ COMMANDS:
	⠙ public	Generates asset bundling for standard public files

	⠙ view	Generates asset bundling with html file

	⠙ static	Generates bundling for general static files 


⡿ HELP:
	Run [command] help

⡿ OTHERS:
	Run 'assetkit printflags' to print all flags of all commands.

⡿ WARNING:
	Uses internal flag package so flags must precede command name. 
	e.g 'assetkit -cmd.flag=4 run'

```


## Contribution

Please let all suggestions, thoughts and complaints come has issues and associated PRs.
