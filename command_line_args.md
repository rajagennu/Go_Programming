# CommandLine args

- Commandline argument in go are available via `os` package. In OS package, we have variable called as `Args`.
- `os.Args` is of slice type
- `os.Args[0]` is the name of the command. 
- from os.Args[1] to os.Args[len(os.Args)] will iterate over all the command line arguments that passed as input to the go program. 
- 