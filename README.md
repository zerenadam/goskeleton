# Goskeleton

Goskeleton is a lightweight CLI tool written in Go for scaffolding Go projects.  
It can create a new project folder with a starter `main.go` and `.gitignore`,  
and optionally initialize a Git repository with a remote.  
It also generates standard Go folders like `internal` and `pkg`.

---

## Features

- Create a new Go project folder structure
- Generate `cmd/<projectname>/main.go` with a starter template
- Generate a `.gitignore` file with `.env` and `.DS_Store`
- Create standard Go folders: `internal` and `pkg`
- Optional Git initialization and remote setup
- Simple CLI usage: `goskeleton <projectname>` or `goskeleton <projectname> -git <gitlink>`
- Emoji-based logs for easy tracking

---

## Installation

You can install Goskeleton globally with Go:

```bash
go install github.com/zerenadam/goskeleton/cmd/goskeleton@latest
```

Make sure \$HOME/go/bin (or \$GOPATH/bin) is in your PATH:

```bash
echo 'export PATH=$HOME/go/bin:$PATH' >> ~/.zshrc
source ~/.zshrc
```

Now you can run `goskeleton` from anywhere.

---

## Usage

### Create a project without Git

```bash
goskeleton MyProject
```

This will:

- Create a folder `MyProject`
- Generate `cmd/MyProject/main.go` with a starter Go template
- Create standard Go folders: `internal` and `pkg`
- Generate a `.gitignore` file with `.env` and `.DS_Store` entries

### Create a project with Git

```bash
goskeleton MyProject -git https://github.com/username/MyProject.git
```

This will additionally:

- Initialize a Git repository inside the project folder
- Add all files
- Commit with an emoji 🎸
- Set the branch to `main`
- Add the remote repository
- Push the initial commit

---

## Example

```bash
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
```

---

## Folder structure created

```
TestApp/
├── cmd/
│   └── TestApp/
│       └── main.go
├── internal/
├── pkg/
└── .gitignore
```

---

## Notes

- The `-git` flag is optional
- The project name is **required**
- The CLI safely handles missing Git links
- Designed for Go developers who want a fast project skeleton

---

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.  
Make sure to follow Go formatting standards (`gofmt`) and keep commits descriptive.

---

## License

MIT License