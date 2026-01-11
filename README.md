# Goskeleton

Goskeleton is a lightweight CLI tool written in Go for scaffolding Go projects.  
It can create a new project folder with a starter `main.go` and `.gitignore`,  
and optionally initialize a Git repository with a remote.

## Features

- Create a new Go project folder structure
- Generate `main.go` with a starter template
- Generate a `.gitignore` file with standard entries
- Create standard Go folders: `internal` and `pkg`
- Optional Git initialization and remote setup
- Simple CLI usage: `goskeleton <projectname>` or `goskeleton <projectname> -git <gitlink>`

## Installation

Install Goskeleton globally with Go:

```bash
go install github.com/zerenadam/goskeleton/cmd/goskeleton@latest```

## Add Goskeleton to your PATH

Make sure `$HOME/go/bin` (or `$GOPATH/bin`) is in your PATH:

```bash
echo 'export PATH=$HOME/go/bin:$PATH' >> ~/.zshrc
source ~/.zshrc
```

## Usage

### Create a project without Git

```bash
goskeleton MyProject
```
This will:

Create a folder MyProject

Generate cmd/MyProject/main.go with a starter Go template

Create standard Go folders: internal and pkg

Generate a .gitignore file with .env and .DS_Store entries

Create a project with Git
goskeleton MyProject -git https://github.com/username/MyProject.git


This will additionally:

Initialize a Git repository inside the project folder

Add all files

Commit with an emoji 🎸

Set the branch to main

Add the remote repository

Push the initial commit

Example
$ goskeleton TestApp -git https://github.com/zerenadam/TestApp.git
Creating project: TestApp
Successfully created cmd/TestApp/main.go
Successfully created internal/
Successfully created pkg/
Successfully created .gitignore
Successfully ran: init
Successfully ran: add .
Successfully ran: commit -m 🎸
Successfully ran: branch -m main
Successfully ran: remote add origin https://github.com/zerenadam/TestApp.git
Successfully ran: push -u origin main

Folder structure created
TestApp/
├── cmd/
│   └── TestApp/
│       └── main.go
├── internal/
├── pkg/
└── .gitignore

Notes

The -git flag is optional

The project name is required

The CLI safely handles missing Git links

Logs include emojis for easy tracking

Contributing

Contributions are welcome! Feel free to open issues or pull requests.

License

MIT License