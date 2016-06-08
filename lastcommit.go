package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
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

// GetCommitMessagePrefix retuns the issue prefix from a commit message
func GetCommitMessagePrefix(msg string) string {
	// ^.{7,} (.+):
	re := regexp.MustCompile("^.{7,} (.+):")
	matches := re.FindStringSubmatch(msg)

	if len(matches) > 0 {
		return matches[1]
	}

	return ""
}

func main() {
	// figure out what path we're in
	// dir := getDirectory()

	// find the last commit
	lastCommitMsg := getLastCommitMessage()
	prefix := GetCommitMessagePrefix(lastCommitMsg)
	fmt.Printf("last commit's prefix: %s\n", prefix)

	// read its prefix

}
