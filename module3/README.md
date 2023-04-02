# 模块三作业

### 1、构建linux运行文件

> 进入main.go文件目录，执行

```shell
GOOS=linux GOARCH=amd64 go build -o httpserver .
```

### 2、编写Dockerfile

```dockerfile
FROM golang:alpine
EXPOSE 80
RUN mkdir /app
WORKDIR /app
COPY httpserver /app/httpserver
ENTRYPOINT ["./httpserver"]
```

### 3、构建docker镜像

```shell
docker build -t httpserver:v1.0 ./
```

### 4、上传镜像

```shell
docker tag httpserver:v1.0 heyu/httpserver:v1.0
docker push heyu/httpserver:v1.0
```

### 5、运行镜像

```shell
docker run --name httpserver -d -p 80:80 heyu/httpserver:v1.0
```

