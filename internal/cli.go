package internal

import (
	"fmt"
	"os"
)

func ParseCli() (projectname, link string) {
	args := os.Args[1:]
	gitlink := ""

	if len(args) < 1 {
		fmt.Println("Usage: goskeleton <projectname> [-git <gitlink>]")
		os.Exit(1)
	}

	for i := 1; i < len(args); i++ {
		if args[i] == "-git" && i+1 < len(args) {
			gitlink = args[i+1]
			i++
		}
	}

	fmt.Println("Project Name:", args[0])
	if gitlink != "" {
		fmt.Println("Git Link:", gitlink)
	}

	return args[0], gitlink

}
