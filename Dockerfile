FROM golang:1.17.2-alpine3.14 as builder

RUN mkdir -p /go/src/httpserver

WORKDIR /go/src/httpserver

# COPY <src> <dest> 其中 <src> 为Dockerfile所在目录的相对路径
COPY /go/net/http/server/* .

# CGO_ENABLED禁用cgo
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/httpserver .

# FROM scratch 不采用 scratch 是因为scratch中无基本工具，不利于问题排查
FROM alpine:3.14.2
WORKDIR /app
COPY --from=builder /go/bin/httpserver /app
EXPOSE 8888
# 
ENTRYPOINT ["/app/httpserver"]
CMD []