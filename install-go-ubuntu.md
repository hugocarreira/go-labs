# go-labs

### Download Go Lang
```  
https://golang.org/dl/  
```  

### Extract  
```
$ tar xvf filename.tar.gz  
```  
### Permissions
```
$ sudo chown -R root:root ./go
$ sudo mv go /usr/local
```  

### Settings GoPath
```  
$ sudo nano ~/.profile  

export GOROOT=/usr/local/go  
export GOPATH=$HOME/work  
export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin  

$ source ~/.profile
```

### Work dir
```
$ mkdir -p $HOME/work/src/github.com/hugocarreira/
```
