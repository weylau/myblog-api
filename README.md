## 安装

1、安装 Go 并且设置好你的 Go 工作空间

2、govendor依赖管理工具
```
$ go get github.com/kardianos/govendor
```
3、创建项目
```$xslt
$ mkdir -p $GOPATH/src/github.com/weylau/myblog-api && cd "$_"
$ govendor init
$ govendor fetch github.com/gin-gonic/gin@v1.3
$ govendor fetch github.com/go-sql-driver/mysql@v1.4.0
$ govendor fetch github.com/jinzhu/gorm@v1.9.8
```
4、main.go
复制模板
https://github.com/gin-gonic/examples/blob/master/basic/main.go

5、启动
```$xslt
$ go run main.go
```