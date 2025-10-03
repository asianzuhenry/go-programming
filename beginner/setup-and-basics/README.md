# Setup and Basics

<!-- Installing Go, GOROOT & GOPATH, Go modules

go run, go build, go fmt, go vet, go test

Project structure & package management -->
This section covers the basics of setting up a Go development environment and introduces fundamental concepts and tools you'll need to get started with Go programming.
## Installing Go
To install Go, follow these steps:
1. Visit the official Go website: [https://golang.org/dl/](https://golang.org/dl/)
2. Download the appropriate installer for your operating system (Windows, macOS, or Linux).
3. Follow the installation instructions provided on the website.
4. Verify the installation by opening a terminal and running:
   ```bash
   go version
   ```
   You should see the installed Go version.
## Setting Up GOROOT and GOPATH
- **GOROOT**: This is the directory where Go is installed. It is usually set automatically during installation. You can check it by running:
  ```bash
  go env GOROOT
  ```
- **GOPATH**: This is the workspace directory where your Go projects and dependencies are stored. By default, it is set to `$HOME/go` on Unix-based systems and `%USERPROFILE%\go` on Windows. You can check it by running:
  ```bash
  go env GOPATH
  ```
You can change the GOPATH by setting the `GOPATH` environment variable in your shell configuration file (e.g., `.bashrc`, `.zshrc`, or `.bash_profile`).
## Go Modules
Go modules are the standard way to manage dependencies in Go projects. To create a new module, follow these steps:
1. Create a new directory for your project and navigate into it:
   ```bash
   mkdir myproject
   cd myproject
   ```
2. Initialize a new Go module:
   ```bash
   go mod init myproject
   ```
3. You can now add dependencies to your project using:
   ```bash
   go get <dependency>
   ```
4. To tidy up your `go.mod` file and remove unused dependencies, run:
   ```bash
   go mod tidy
   ```
## Common Go Commands
- **Run a Go program**:
  ```bash
  go run main.go
  ```
- **Build a Go program**:
  ```bash
  go build -o myprogram main.go
  ```
- **Format Go code**:
  ```bash
  go fmt ./...
  ```
- **Vet Go code for potential issues**:
  ```bash
  go vet ./...
  ```
- **Test Go code**:
  ```bash
  go test ./...
  ```
## Project Structure and Package Management
A typical Go project structure looks like this:
```myproject/
├── go.mod
├── go.sum
├── main.go
├── pkg/
│   └── mypackage/
│       └── myfile.go
└── cmd/
    └── mycommand/
        └── main.go
```
- `go.mod`: Contains module information and dependencies.
- `go.sum`: Contains checksums for module dependencies.
- `main.go`: The entry point of your application.
- `pkg/`: Contains reusable packages.
- `cmd/`: Contains application entry points for different commands.
By following this guide, you should have a solid foundation for setting up a Go development environment and understanding the basics of Go programming. Happy coding!