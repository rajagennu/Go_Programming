# Build process

- you must create a mod to build your go project. here is terminal trace 

```bash
➜  go-ch1 go mod init github.com/rajagennu/ch1-go
go: creating new go.mod: module github.com/rajagennu/ch1-go
go: to add module requirements and sums:
	go mod tidy
➜  go-ch1 go mod tidy
➜  go-ch1 make                                   
go fmt ./...
hello.go
go vet ./...
go build hello.go
➜  go-ch1 ls
go.mod  hello  hello.go  Makefile
➜  go-ch1 ./hello 
Hello Go!
➜  go-ch1 

```

- here is my make file

```make
.DEFAULT_GOAL := build

fmt:
		go fmt ./...
.PHONY:fmt

lint: fmt
		golint ./...
.PHONY:lint

vet: fmt
		go vet ./...
.PHONY:vet

build: vet
		go build hello.go
.PHONY:build

```

- other commands as per above definition
```
make fmt
make lint
make vet
```
