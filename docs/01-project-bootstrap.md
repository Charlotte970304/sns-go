# Project Bootstrap – Getting Started with Go

## 1. Why Go
I chose Go as the backend language for this project because:
- It is designed for backend and server-side development
- It has strong support for concurrency, which is important for SNS systems
- Its simplicity and explicit design encourage clear system boundaries

This project focuses on backend architecture rather than rapid prototyping,
which aligns well with Go’s strengths.

---

## 2. Installing Go
Go was downloaded from the official website:

https://go.dev/dl

After installation, the setup was verified by running:

```bash
go version
```
Seeing the Go version confirmed that the toolchain was installed correctly.

---

## 3. What Is a Go Module

A Go module represents a project-level boundary.

By running:

```bash
go mod init sns-go
```

the project gained:
- A clear project identity
- Dependency and version management
- A defined root for internal imports

Using Go modules allows the project to scale cleanly as more packages
and dependencies are added.

---

## 4. First Runnable Program

A minimal executable entry point was created at:

```bash
cmd/api/main.go
```
```go
package main

import "fmt"

func main() {
    fmt.Println("SNS backend bootstrapped")
}
```
The program was successfully executed using:
```bash
go run cmd/api/main.go
```
This confirmed that the project structure and Go module
were set up correctly.