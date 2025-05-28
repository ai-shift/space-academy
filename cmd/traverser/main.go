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

type ExtInfo struct {
	Count int   `json:"count"`
	Size  int64 `json:"size"`
}

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
	exts := make(map[string]*ExtInfo)
	err = filepath.WalkDir(traversePath, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
		fi, err := os.Stat(path)
		if err != nil {
			return err
		}
		ext := filepath.Ext(path)
		info, ok := exts[ext]
		if !ok {
			info = &ExtInfo{}
			exts[ext] = info
		}
		info.Count += 1
		info.Size += fi.Size()
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
		return exts[keys[i]].Size > exts[keys[j]].Size
	})

	log.Println("Printing results as a table")
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	fmt.Fprintln(w, "File extension\tCount\tSize")
	for _, key := range keys {
		fmt.Fprintf(w, "%s\t%d\t%d\n", key, exts[key].Count, exts[key].Size)
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
