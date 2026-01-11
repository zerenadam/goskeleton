package internal

import (
	"log"
	"os/exec"
	"strings"
)

func RunGitCommands(dir, gitlink string) {
	commands := [][]string{
		{"init"},
		{"add", "."},
		{"commit", "-m", "🦴"},
		{"branch", "-m", "main"},
		{"remote", "add", "origin", gitlink},
		{"push", "-u", "origin", "main"},
	}

	for _, commandArg := range commands {
		cmd := exec.Command("git", commandArg...)
		cmd.Dir = dir

		if output, err := cmd.CombinedOutput(); err != nil {
			log.Fatal("Error Running", strings.Join(commandArg, " "), ":", err, string(output))
		} else {
			log.Println("Succesfully Ran : ", strings.Join(commandArg, " "))
		}
	}
}
