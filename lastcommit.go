package main

import (
	"fmt"
	"os"
)

func getDirectory() string {
	wd, _ := os.Getwd()
	return wd
}

func main() {
	// figure out what path we're in
	// find the last commit
	// read its prefix
	dir := getDirectory()

	fmt.Println("you are in", dir)
}
