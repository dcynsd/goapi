FROM golang:1.21-alpine as builder

LABEL authors="dcynsd"
LABEL email="dcynsd@gmail.com"

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /app

COPY . .

RUN go build -o goapi .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/goapi .
COPY --from=builder /app/config ./config

# 指定运行时环境变量
ENV GIN_MODE=release \
    APP_PORT=80

EXPOSE 80

ENTRYPOINT ["./goapi"]
