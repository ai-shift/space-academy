package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sort"
	"text/tabwriter"
)

const (
	resultPath = "result.json"
)

func main() {
	// Parse CLI flags
	var traversePath string
	flag.StringVar(&traversePath, "path", ".", "Root path to traverse files")
	flag.Parse()
	traversePath, err := filepath.Abs(traversePath)
	if err != nil {
		log.Fatal("Failed to resolve path with", err)
	}

	log.Println("Traversing", traversePath)
	exts := make(map[string]int)
	err = filepath.WalkDir(traversePath, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
		exts[filepath.Ext(path)] += 1
		return nil
	})
	if err != nil {
		log.Fatal("Failed to traverse with", err)
	}

	log.Println("Sorting results")
	keys := make([]string, 0, len(exts))
	for key := range exts {
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return exts[keys[i]] > exts[keys[j]]
	})

	log.Println("Printing results as a table")
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	fmt.Println("File extension\tAmount")
	for _, key := range keys {
		fmt.Fprintf(w, "%s\t%d\n", key, exts[key])
	}
	w.Flush()

	log.Println("Saving results to", resultPath)
	extsB, _ := json.Marshal(exts)
	err = os.WriteFile(resultPath, extsB, 0644)
	if err != nil {
		log.Fatal("Failed to write result JSON file with", err)
	}

	log.Println("Success")
}
