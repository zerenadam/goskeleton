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
	)

	files := map[string]string{
		filepath.Join(projectname, "cmd", projectname, "main.go"): maingo,
		filepath.Join(projectname, ".gitignore"):                  gitignore,
	}

	for path, content := range files {
		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			log.Fatal(err)
		}
		log.Println("Succesfully Created : ", path)
	}
}
