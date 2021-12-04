# Go Lint & Vet


Installation:

```bash
go install golang.org/x/lint/golint@latest
```

And run it with

```bash
golint ./..
go vet ./..
```

We can also use `golangci-lint` for doing all the lint, vet checks

Installation

```bash
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.43.0

golangci-lint --version

```

configuration file : `.golangci.yml`
location: root directory of the project



