package main

import (
	"github.com/zerenadam/goskeleton/internal"
)

func main() {
	projectname, gitlink, gitflag := internal.ParseCli()

	internal.CreateFolders(projectname)
	internal.CreateFiles(projectname)
	if gitflag == true {
		internal.RunGitCommands(projectname, gitlink)
	}

}
