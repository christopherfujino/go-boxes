package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	fmt.Printf("Found root at %s\n", root)

	var projectRoots = findAllRecursive("go.mod", root)

	var errors = 0

	for _, root := range projectRoots {
		var cmd = exec.Command("go", "vet", ".")
		cmd.Dir = root
		fmt.Printf("About to vet %s...\n", root)
		var stderrBuffer = strings.Builder{}
		cmd.Stderr = &stderrBuffer
		err := cmd.Run()
		if err != nil {
			fmt.Fprintf(os.Stderr, `Error

%s

%s
`, stderrBuffer.String(), err.Error())
			errors += 1
		}
	}

	if errors > 0 {
		panic("Failure")
	}
}
