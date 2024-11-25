package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Options struct {
	maxDepth      int
	showHidden    bool
	excludeFilter string
	showTimer     bool
}

func main() {
	maxDepth := flag.Int("depth", -1, "Maximum depth to traverse (-1 for unlimited)")
	showHidden := flag.Bool("hidden", false, "Show hidden files and directories")
	excludeFilter := flag.String("exclude", "", "Comma-separated list of names to exclude")
	showTimer := flag.Bool("timer", false, "Show execution time")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS] [DIRECTORY]\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nIf no directory is specified, the current directory will be used.\n")
	}

	flag.Parse()

	root := "."
	if args := flag.Args(); len(args) > 0 {
		root = args[0]
	}

	opts := Options{
		maxDepth:      *maxDepth,
		showHidden:    *showHidden,
		excludeFilter: *excludeFilter,
		showTimer:     *showTimer,
	}

	startTime := time.Now()

	fmt.Println(root)
	printTree(root, "", 0, opts)

	if opts.showTimer {
		duration := time.Since(startTime)
		fmt.Printf("\nExecution time: %v\n", duration)
	}
}

// Basic DFS traversal of the file system
func printTree(path string, prefix string, depth int, opts Options) error {
	if opts.maxDepth != -1 && depth >= opts.maxDepth {
		return nil
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	excludeList := strings.Split(opts.excludeFilter, ",")

	for i, entry := range entries {
		name := entry.Name()

		if !opts.showHidden && strings.HasPrefix(name, ".") {
			continue
		}

		// Skip excluded files/directories
		excluded := false
		for _, exclude := range excludeList {
			if exclude != "" && name == exclude {
				excluded = true
				break
			}
		}
		if excluded {
			continue
		}

		isLast := i == len(entries)-1

		newPrefix := prefix
		if isLast {
			fmt.Printf("%s└── %s\n", prefix, name)
			newPrefix += "    "
		} else {
			fmt.Printf("%s├── %s\n", prefix, name)
			newPrefix += "│   "
		}

		if entry.IsDir() {
			printTree(filepath.Join(path, name), newPrefix, depth+1, opts)
		}
	}

	return nil
}
