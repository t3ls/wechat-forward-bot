FROM golang:1.21-alpine as builder

RUN apk --no-cache add git && \
    export GOPROXY=https://goproxy.cn,direct

COPY . /root/build

WORKDIR /root/build

RUN go mod download && go build -o server main.go

FROM alpine:latest as prod

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=0 /root/build/server .

CMD ["./server"]