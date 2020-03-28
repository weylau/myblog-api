#源镜像
FROM golang:1.13

#设置工作目录
WORKDIR $GOPATH/src

#拉取项目代码
RUN git clone https://github.com/weylau/myblog-api.git
#切换工作目录
WORKDIR $GOPATH/src/myblog-api

#设置代理
ENV GOPROXY https://goproxy.io

#go构建可执行文件
RUN go build -o myblog-api .

#暴露端口
EXPOSE 9001

#最终运行docker的命令
ENTRYPOINT  ["nohup","./myblog-api","&"]