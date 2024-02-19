FROM golang:1.18

WORKDIR /go-pro/mall

COPY . .

CMD ["/go-pro/mall/main"]
