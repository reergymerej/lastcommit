package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func getDirectory() string {
	wd, _ := os.Getwd()
	return wd
}

func getLastCommitMessage() string {
	out, err := exec.Command("git", "log", "--oneline").Output()
	if err != nil {
		log.Fatal(err)
	}

	return string(out[:])
}

func getCommitMessagePrefix(msg string) string {
	return msg
}

func main() {
	// figure out what path we're in
	// dir := getDirectory()

	// find the last commit
	lastCommitMsg := getLastCommitMessage()
	prefix := getCommitMessagePrefix(lastCommitMsg)
	fmt.Printf("The last commit's prefix was: %s\n", prefix)

	// read its prefix

}
