package main

import (
	"fmt"
	"os"
	"path/filepath"
	"flag"

	"github.com/docker/docker/builder/dockerignore"
	"github.com/docker/docker/pkg/fileutils"
)

var excludes []string

func main() {
	flag.Parse()
	root := flag.Arg(0)
	if len(root) == 0 {
		fmt.Printf ("Pass in file path to examine.\n")
		panic("Missing main argument")
	}

	fmt.Println("Will apply .dockerignore on path", root)
	file, dockerignoreerr := os.Open(".dockerignore")
	check(dockerignoreerr)

	excludes, _ = dockerignore.ReadAll(file)
	fmt.Println("Patterns are",excludes)

	filepath.Walk(root, visit)
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func visit(file string, f os.FileInfo, err error) error {
	rm, _ := fileutils.Matches(file, excludes)
	if rm {
		fmt.Printf("Removing file %v\n" , file)
		var err = os.Remove(file)
		check(err)
	}
	return nil
}
