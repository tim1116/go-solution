FROM golang:alpine AS builder

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

# 将代码复制到容器中
COPY . .
# 将我们的代码编译成二进制可执行文件app
RUN go build -o app .

# 分阶段构建
FROM scratch

COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai


COPY ./template /template
COPY ./public /public
COPY ./config /config


COPY --from=builder /build/app /

# 需要运行的命令
ENTRYPOINT ["/app"]