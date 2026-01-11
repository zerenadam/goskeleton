package internal

import (
	"log"
	"os"
	"path/filepath"
)

func CreateFiles(projectname string) {
	const (
		maingo = `package main
		
import (
	"fmt"
)

func main(){
	fmt.Println("Hello World")
	}`
		gitignore = `.env
.DS_Store`

		readme = `
---

[![made with goskeleton](https://img.shields.io/badge/made%20with-goskeleton%20🦴-blue)](https://github.com/zerenadam/goskeleton)`
	)

	files := map[string]string{
		filepath.Join(projectname, "cmd", projectname, "main.go"): maingo,
		filepath.Join(projectname, ".gitignore"):                  gitignore,
		filepath.Join(projectname, "README.md"):                   readme,
	}

	for path, content := range files {
		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			log.Fatal(err)
		}
		log.Println("Succesfully Created : ", path)
	}
}
