### go_install.md

- install with repository location as below
```bash
go install github.com/rakyll/hey@latest
```

- `latest` refer to latest veriosn, replace it with version number to get specifc version

- once downloaded, go install will take care of compiling, installing automatically.

- installation location " $GOPATH/bin"

- to get $GOPATH location `go env | grep GOPATH`

```bash
➜  ch1 git:(main) ✗ go env | grep GOPATH
GOPATH="/home/agastya/go"                                                  
➜  ~ cd /home/agastya/go/bin                                               
➜  bin ls
hey                                                                        
➜  bin ./hey google.com

Summary:                                                                   
  Total:        0.0017 secs        
```