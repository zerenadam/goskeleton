package internal

import (
	"log"
	"os"
	"path/filepath"
)

func CreateFolders(projectname string) {
	dirs := []string{
		filepath.Join(projectname, "cmd", projectname),
		filepath.Join(projectname, "internal"),
		filepath.Join(projectname, "pkg"),
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			log.Fatal("Error creating", dir, ":", err)
		}
		log.Println("Succesfully created : ", dir)
	}
}
