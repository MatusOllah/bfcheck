package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/fatih/color"
	"github.com/jessevdk/go-flags"
	"github.com/ztrue/tracerr"
)

const version = "1.1.0"

type Config struct {
	PosArgs struct {
		Path string `description:"Path to FNF mod" default:"."`
	} `positional-args:"yes" required:"yes"`
	Verbose bool `short:"v" long:"verbose" description:"Show verbose debug information"`
	//Path        string `short:"p" long:"path" description:"Path to FNF mod" default:"."`
	Color       bool `short:"c" long:"color" description:"Color output"`
	ShowLines   bool `short:"l" long:"show-lines" description:"Show lines when printing found instances"`
	WriteReport bool `short:"r" long:"write-report" description:"Write a report in JSON format"`
	Version     bool `long:"version" description:"Print version and exit"`
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

	r, err := checkDir(cfg.PosArgs.Path)
	if err != nil {
		tracerr.Print(err)
		os.Exit(1)
	}

	if r.NumInstances != 0 {
		if cfg.Color {
			color.New(color.Bold).Printf("instances: ")
			fmt.Println(r.NumInstances)
		} else {
			fmt.Printf("instances: %d\n", r.NumInstances)
		}
	} else if r.NumInstances == 0 {
		if cfg.Color {
			color.Green("ALL OK")
		} else {
			fmt.Println("ALL OK")
		}
	}

	if cfg.WriteReport {
		if err := WriteReport(r); err != nil {
			tracerr.Print(err)
			os.Exit(1)
		}
	}
}

func WriteReport(r *Report) error {
	path := fmt.Sprintf("bfcheck_report_%d.json", r.Time)

	if cfg.Verbose {
		if cfg.Color {
			color.New(color.FgCyan).Printf("Writing report to ")
			fmt.Println(path)
		} else {
			fmt.Printf("Writing report to %s\n", path)
		}
	}

	b, err := r.Encode()
	if err != nil {
		return tracerr.Wrap(err)
	}

	f, err := os.Create(path)
	if err != nil {
		return tracerr.Wrap(err)
	}

	written, err := f.Write(b)
	if err != nil {
		return tracerr.Wrap(err)
	}

	if cfg.Verbose {
		if cfg.Color {
			color.New(color.FgGreen, color.Bold).Printf("Wrote %d bytes\n", written)
		} else {
			fmt.Printf("Wrote %d bytes\n", written)
		}
	}

	f.Close()

	return nil
}
