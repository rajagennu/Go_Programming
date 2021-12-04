### GOPATH Environment variable

- 3rd party go libraries installed via `go install` will be stored at $HOME/go folder.
- source code of these libraries stored at $HOME/go/src
- binaries will be stored at $HOME/go/bin
- so actually where this 3rd party go libs can be installed also configurable by $GOPATH environment variable.
- defining $GOPATH means, definiing your Go workspace.
- By Adding $GOPATH/bin to your $PATH, makes the binaries inside available to execute at command line path.
```bash
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```
- make sure adding those lines in your `/.bashrc` or in your `~/.bash_profile` or `~/.zshrc` ( simply putting - your default shell)
- all the environment variables used by go can be listed using `go env` command.

