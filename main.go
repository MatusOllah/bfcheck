package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/fatih/color"
	"github.com/jessevdk/go-flags"
	"github.com/ztrue/tracerr"
)

const version = "1.0.0"

type Config struct {
	Verbose     bool   `short:"v" long:"verbose" description:"Show verbose debug information"`
	Path        string `short:"p" long:"path" description:"Path to FNF mod" default:"."`
	Color       bool   `short:"c" long:"color" description:"Color output"`
	ShowLines   bool   `short:"l" long:"show-lines" description:"Show lines when printing found instances"`
	WriteReport bool   `short:"r" long:"write-report" description:"Write a report in JSON format"`
	Version     bool   `long:"version" description:"Print version and exit"`
}

var cfg Config

func main() {
	if _, err := flags.NewParser(&cfg, flags.HelpFlag|flags.PassDoubleDash).Parse(); err != nil {
		tracerr.Print(err)
		os.Exit(1)
	}

	c := color.New(color.FgCyan, color.Bold).SprintfFunc()
	w := color.New(color.FgWhite, color.Bold).SprintfFunc()

	if cfg.Version {
		if cfg.Color {
			fmt.Fprintln(color.Output, c("bfcheck"), w("version"), version)
			fmt.Fprintln(color.Output, c("Go"), w("version"), runtime.Version())
		} else {
			fmt.Println("bfcheck version", version)
			fmt.Println("Go version", runtime.Version())
		}

		os.Exit(0)
	}

	num, err := checkDir(cfg.Path)
	if err != nil {
		tracerr.Print(err)
		os.Exit(1)
	}

	if num != 0 {
		if cfg.Color {
			color.New(color.Bold).Printf("instances: ")
			fmt.Println(num)
		} else {
			fmt.Printf("instances: %d\n", num)
		}
	} else if num == 0 {
		if cfg.Color {
			color.Green("ALL OK")
		} else {
			fmt.Println("ALL OK")
		}
	}
}
