package internal

import (
	"log"
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
			log.Fatal("Error Creating", file, ":", err)
		}
		log.Println("Succesfully Created : ", file)
	}
}
