package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/ztrue/tracerr"
)

func checkDir(pathToFNF string) (*Report, error) {
	if cfg.Verbose {
		if cfg.Color {
			color.New(color.FgCyan).Printf("checking ")
			fmt.Println(pathToFNF)
		} else {
			fmt.Printf("checking %s\n", pathToFNF)
		}
	}

	r := &Report{
		Time:      time.Now().UnixNano(),
		Path:      pathToFNF,
		Instances: []Instance{},
	}

	err := filepath.Walk(filepath.Join(pathToFNF, "assets"), func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return tracerr.Wrap(err)
		}

		if !info.IsDir() {
			if filepath.Ext(path) != ".txt" &&
				filepath.Ext(path) != ".md" &&
				filepath.Ext(path) != ".rst" &&
				filepath.Ext(path) != ".json" &&
				filepath.Ext(path) != ".toml" &&
				filepath.Ext(path) != ".yaml" &&
				filepath.Ext(path) != ".xml" &&
				filepath.Ext(path) != ".yml" &&
				filepath.Ext(path) != ".hscript" &&
				filepath.Ext(path) != ".lua" &&
				filepath.Ext(path) != ".hx" &&
				filepath.Ext(path) != ".c" &&
				filepath.Ext(path) != ".h" &&
				filepath.Ext(path) != ".cpp" &&
				filepath.Ext(path) != ".hpp" &&
				filepath.Ext(path) != ".cxx" &&
				filepath.Ext(path) != ".hxx" &&
				filepath.Ext(path) != ".go" {
				return nil
			}

			fileInstances, err := checkFile(path)
			if err != nil {
				return tracerr.Wrap(err)
			}

			r.Instances = append(r.Instances, fileInstances...)
		}

		return nil
	})
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	r.NumInstances = len(r.Instances)

	return r, nil
}

func checkFile(path string) ([]Instance, error) {
	if cfg.Verbose {
		if cfg.Color {
			color.New(color.FgCyan).Printf("checking ")
			fmt.Println(path)
		} else {
			fmt.Printf("checking %s\n", path)
		}
	}

	instances := []Instance{}

	file, err := os.Open(path)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	sc := bufio.NewScanner(file)
	for ln := 1; sc.Scan(); ln++ {
		for _, pattern := range blacklist {
			if strings.Contains(strings.ToLower(sc.Text()), strings.ToLower(pattern)) {
				col := strings.Index(strings.ToLower(sc.Text()), strings.ToLower(pattern))

				fmt.Println()
				if cfg.Color {
					color.New(color.Bold).Printf("%s:%d:%d: found \"%s\"\n", path, ln, col+1, pattern)
				} else {
					fmt.Printf("%s:%d:%d: found \"%s\"\n", path, ln, col+1, pattern)
				}

				if cfg.ShowLines {
					fmt.Printf("\t%s\n", sc.Text())

					printArrows(len(sc.Text())+(col-len(sc.Text())), len(pattern))

				}

				fmt.Println()

				instances = append(instances, Instance{
					File:   path,
					Line:   ln,
					Column: col,
				})
			}
		}
	}

	return instances, nil
}

func printArrows(index, length int) {
	fmt.Printf("\t")
	for i := 0; i < index; i++ {
		fmt.Printf(" ")
	}
	for i := 0; i < length; i++ {
		if cfg.Color {
			color.New(color.FgGreen).Printf("^")
		} else {
			fmt.Printf("^")
		}
	}
	fmt.Println()
}
