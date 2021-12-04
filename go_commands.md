### Go Commands
- `go run file_name.go`
- `go build file_name.go`
- `go build -o my_binary_custom_name file_name.go`
- `go install github.com/rakyll/hey@latest`
	- latest refers latest version, you can replace with specific version
	- binary will be stored at $GOPATH/bin location.
	- go install will also covers the dependencies  required
- `go fmt filename.go` will format source code automatically.
- `goimports -l -w .`  (goimports)[goimports.md]
