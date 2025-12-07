
`sudo tar -zxvf go$VERSION.$OS-$ARCH.tar.gz -C /usr/local`     


`sudo vim /etc/profile`    

```
export GOFOLDER=go1-15-11
export GOROOT=/usr/local/$GOFOLDER
export GOPATH=/var/${GOFOLDER}-path
export PATH=$PATH:${GOROOT}/bin:${GOPATH}/bin

export GOPROXY=https://goproxy.io,direct
export GOPRIVATE=gitlab.com,gitee.com,gitea.com
```


`source /etc/profile`  

`source ~/.bashrc`  
