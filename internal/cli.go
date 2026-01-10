package internal

import (
	"flag"
	"fmt"
	"os"
)

func ParseCli() (projectname, gitlink string) {
	gitLink := flag.String("git", "", "Git remote URL (optional)")
	flag.Parse()

	args := flag.Args()

	if len(args) < 1 {
		fmt.Println("Usage: goskeleton <projectname> [-git <gitlink>]")
		os.Exit(1)
	}

	return args[0], *gitLink

}
