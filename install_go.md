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

