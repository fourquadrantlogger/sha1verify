# sha1verify
sha1 hash
## command

```
go build
./sha1verify
```

或者
```
go build
./sha1verify  --folder="/MacData/GOPATH/src/code.aliyun.com/timeloveboy/sha1verify/"  --outfile="my.txt" --gitignore=".gitignore"
```

```
go run sha1verify.go
go run 
```

## 过滤配置文件
仅支持过滤文件，和文件夹

```
/out
/.gitignore
/.idea
/.git
```