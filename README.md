# 🦴 goskeleton

a lightweight go project scaffolding tool that generates a clean project structure and optionally initializes git with your first commit

## ✨ features

- 🚀 instant project setup with clean folder structure
- 📁 automatic directory creation
- 📝 generates starter files
- 🎸 optional git initialization and first commit
- 🔗 automatic remote repository linking
- ⚡ simple and fast no configuration needed

## 📥 installation

### via go install

```bash
go install github.com/zerenadam/goskeleton@latest
```

### via go run

```bash
go run github.com/zerenadam/goskeleton <projectname>
```

## 🎯 usage

### basic usage

create a new project without git initialization

```bash
goskeleton myproject
```

### with git initialization

create a project and automatically initialize git add files commit and push to remote

```bash
goskeleton myproject -git https://github.com/username/myproject.git
```

this will automatically

- initialize a git repository
- add all files
- create initial commit with 🎸 emoji
- rename branch to main
- add remote origin
- push to remote repository

## 📁 generated folder structure

```
myproject/
├── cmd/
│   └── myproject/
│       └── main.go          # application entry point with hello world
├── internal/                # private application code
├── pkg/                     # public libraries
└── .gitignore              # ignores .env and .DS_Store
```

### folder breakdown

**cmd/projectname/** contains your main application entry point

**internal/** for private application code that shouldnt be imported by other projects

**pkg/** for code that can be shared and imported by external projects

**.gitignore** preconfigured to ignore environment files and macos system files

## 🚀 quick start example

### create a basic project

```bash
goskeleton awesomeapp
cd awesomeapp
go run cmd/awesomeapp/main.go
```

output
```
Hello World
```

### create project with git

first create an empty repository on github then

```bash
goskeleton awesomeapp -git https://github.com/yourusername/awesomeapp.git
```

goskeleton will create the structure initialize git and push everything to your remote repository automatically

## 📝 generated main.go

the starter main.go file includes a simple hello world program

```go
package main

import (
    "fmt"
)

func main(){
    fmt.Println("Hello World")
}
```

## 🎸 git workflow

when using the `-git` flag goskeleton executes these commands automatically

```bash
git init                                    # initialize repository
git add .                                   # stage all files
git commit -m "🎸"                         # create initial commit
git branch -m main                          # rename to main branch
git remote add origin <your-repo-url>       # add remote
git push -u origin main                     # push to remote
```

## 📋 prerequisites

- go 118 or higher
- git installed if using the `-git` flag

## 💡 usage tips

### after project creation

navigate to your project and start building

```bash
cd myproject
go mod init github.com/username/myproject
go run cmd/myproject/main.go
```

### organizing your code

use the folder structure as a foundation

- put your main application in `cmd/projectname/`
- add internal packages in `internal/`
- create shared libraries in `pkg/`

### example structure as your project grows

```
myproject/
├── cmd/
│   └── myproject/
│       └── main.go
├── internal/
│   ├── handlers/
│   ├── models/
│   └── services/
├── pkg/
│   └── utils/
└── .gitignore
```

## 🔧 command line options

```bash
goskeleton <projectname>                    # create project only
goskeleton <projectname> -git <repo-url>    # create project and initialize git
```

### examples

```bash
goskeleton myapi
goskeleton webserver -git https://github.com/user/webserver.git
goskeleton tool -git git@github.com:user/tool.git
```

## ⚠️ notes

- the project name will be used as the folder name and subfolder under cmd
- make sure your git repository exists and is empty before using the `-git` flag
- the initial commit message is 🎸 guitar emoji
- `.gitignore` is preconfigured to ignore `.env` and `.DS_Store` files

## 🤝 contributing

contributions are welcome feel free to open issues or submit pull requests on github

## 📄 license

check the repository for license information

## 🙏 acknowledgments

built to simplify go project initialization