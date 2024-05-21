package main

import (
	"os"
	"path/filepath"
)

var root = (func() string {
	var cwd, err = os.Getwd()
	if err != nil {
		panic(err)
	}

	return walkUp(cwd)
})()

func walkUp(start string) string {
	entities, err := os.ReadDir(start)
	if err != nil {
		panic(err)
	}

	for _, entity := range entities {
		if entity.IsDir() {
			if entity.Name() == ".git" {
				return start
			}
		}
	}

	return walkUp(filepath.Dir(start))
}

func findAllRecursive(pattern string, root string) []string {
	children, err := os.ReadDir(root)
	if err != nil {
		panic(err)
	}

	var found = []string{}

	for _, child := range children {
		if child.IsDir() {
			found = append(
				found,
				findAllRecursive(pattern, filepath.Join(root, child.Name()))...,
			)
		} else {
			if child.Name() == pattern {
				found = append(found, root)
			}
		}
	}

	return found
}
