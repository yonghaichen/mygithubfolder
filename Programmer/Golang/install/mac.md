 
`tar -C /usr/local -zxvf go$VERSION.$OS-$ARCH.tar.gz`     


`vi ~/.zshrc`   

```
export GOFOLDER=go1-20-1
export GOROOT=/usr/local/$GOFOLDER
export GOPATH=${HOME}/${GOFOLDER}-path
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin

export GOPROXY=https://goproxy.io,direct
export GOPRIVATE=gitlab.com,gitee.com,gitea.com
```


`source ~/.zshrc`   

`export GOPROXY=https://goproxy.cn`