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

	cwd, err := os.Getwd()
	check(err)
	fmt.Printf("Will apply .dockerignore on path %v ( %v )\n", root, cwd)
	file, dockerignoreerr := os.Open(".dockerignore")
	check(dockerignoreerr)

	readDockerIgnore(file)

	filepath.Walk(root, visit)
}

func check(err error) {
    if err != nil {
        panic(err)
    }
}

func readDockerIgnore(file *os.File) {
	var err error
	excludes, err = dockerignore.ReadAll(file)
	check(err)
	fmt.Printf("Patterns are %v\n", excludes)
}

func isDirectory(path string) string {
	fileInfo, err := os.Stat(path)
	check(err)
	if fileInfo.IsDir() {
		return "directory"
	} else {
		return "file"
	}
}

func visit(path string, f os.FileInfo, err error) error {
	rm, err := fileutils.Matches(path, excludes)
	check(err)
	if rm {
		fmt.Printf("Removing %v %v\n", isDirectory(path), path)
		var err = os.Remove(path)
		check(err)
	}
	return nil
}
