package internal

import (
	"fmt"
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
			fmt.Println("Error creating", dir, ":", err)
		}
		fmt.Println("Succesfully created : ", dir)
	}
}
