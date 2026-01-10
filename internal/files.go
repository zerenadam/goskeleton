package internal

import (
	"fmt"
	"os"
	"path/filepath"
)

func CreateFiles(projectname string) {
	files := []string{
		filepath.Join(projectname, "cmd", projectname, "main.go"),
		filepath.Join(projectname, ".gitignore"),
	}

	for _, file := range files {
		if _, err := os.Create(file); err != nil {
			fmt.Println("Error Creating", file, ":", err)
		}
		fmt.Println("Succesfully Created : ", file)
	}
}
