### Go Installation in Linux
-> Go to golang.org/dl
-> Download latest tar
-> from the execution directory 
```bash
sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.17.4.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
```

-> Verify installation with below command
```bash
go version
```


### Go second installation

- if you want to test your program against other versions of go just to test backward compatability through go promised backward comptability

```bash
go get golang.org/dl/go.1.15.6
go.1.15.6 download

# build & run
go.1.15.6 run
go.1.15.6 build

```