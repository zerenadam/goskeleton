package internal

import (
	"log"
	"os/exec"
)

func RunGitCommands(dir, gitlink string) {
	commands := [][]string{
		{"init"},
		{"add", "."},
		{"commit", "-m", "🎸"},
		{"branch", "-m", "main"},
		{"remote", "add", "origin", gitlink},
		{"push", "-u", "origin", "main"},
	}

	for i, commandArg := range commands {
		cmd := exec.Command("git", commandArg...)
		cmd.Dir = dir

		if output, err := cmd.CombinedOutput(); err != nil {
			log.Fatal("Error Running", commands[i], ":", err, string(output))
		} else {
			log.Println("Succesfully Ran : ", commands[i])
		}
	}
}
