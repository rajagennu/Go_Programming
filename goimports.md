### GO Imports

- goimports will help with cleanup import statements, sort the import statements in alphabetical order, cleanup unusued imports, inserting used imports automatically.
- there were some situations where it didnt inserted new import statements automatically, please check your file if automatically imports were added or not.

#### Installation

```bash
go install golang.org/x/tools/cmd/goimports@latest
```

- Usage
```bash
goimports -l -w .

// -l : list files with incorrect format in the console.
// -w : modify the files in place.
// . : in the current working directory

➜  Go_Programming git:(main) ✗ /home/agastya/go/bin/goimports -l -w .  
ch1/hello.go                                                           
➜  Go_Programming git:(main) ✗ 

```
