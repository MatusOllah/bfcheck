package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/jessevdk/go-flags"
	"github.com/ztrue/tracerr"
)

type Config struct {
	Verbose   bool   `short:"v" long:"verbose" description:"Show verbose debug information"`
	Path      string `short:"p" long:"path" description:"Path to FNF mod" default:"."`
	Color     bool   `short:"c" long:"color" description:"Color output"`
	ShowLines bool   `short:"l" long:"show-lines" description:"Show lines when printing found instances"`
}

var cfg Config

func main() {
	if _, err := flags.NewParser(&cfg, flags.HelpFlag|flags.PassDoubleDash).Parse(); err != nil {
		tracerr.Print(err)
		os.Exit(1)
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
