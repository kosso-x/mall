# FROM golang:1.18
FROM golang:1.20

ENV GOPROXY https://goproxy.cn,direct
WORKDIR /go-pro/mall
COPY . .

CMD ["go run /go-pro/mall/main.go"]
