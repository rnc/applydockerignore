package main

import (
	"fmt"
	"os"
	"io"
        "io/ioutil"
        "testing"
)

func exists(name string) bool {
    _, err := os.Stat(name)
    return !os.IsNotExist(err)
}

func TestVisit(t *testing.T) {

	// Setup list of files and patterns to test
	files := []string { "deleteMe", "fileTwo", "fileThree" }
	ignoreString := files[0] + "\n" + "!" + files[1] + "\n" + "fileTh*\n"

	// Temp dir setup and write out the files to test.
        name, err := ioutil.TempDir("", "tempDir")
        check(err)
	fmt.Printf("Created temporary directories in path %v\n", name)
	defer os.RemoveAll (name)
	os.Chdir (name)
	for _, s := range files {
		err = ioutil.WriteFile(s, nil, 0644)
		check(err)
	}

	// Write out a .dockerignore file.
	dfile, err := os.Create(".dockerignore")
	check(err)
	_, err = io.WriteString(dfile, ignoreString)
        check(err)


	// Now the testing...
	//
	// Verify both readDockerIgnore and visit.
	//
	dfile, err = os.Open(".dockerignore")
	readDockerIgnore(dfile)

	if len(excludes) == 0 {
		t.Errorf ("Failed to read .dockerignore")
	}

	for _, s := range files {
		visit (s, nil, nil)
	}

	if exists("deleteMe") || exists ("fileThree") {
		t.Errorf ("File deleteMe still exists")
	}
	if ! exists ("fileTwo") {
		t.Errorf ("File fileTwo does not exist")
	}
}
