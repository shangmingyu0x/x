FROM golang:1.17.2-buster as builder
RUN mkdir -p /build
WORKDIR /build
COPY go/net/http/server/* .
RUN GOOS=linux go build -o /bin/httpserver .

FROM ubuntu
COPY --from=builder /bin/httpserver /bin/httpserver
EXPOSE 8888
ENTRYPOINT ["/bin/httpserver"]
CMD []