FROM golang:1.20 as builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn


WORKDIR /go

RUN git clone http://xiaoe:zxc940429@ah.qyyh.net:3000/xiaoe/qyyh-go.git

WORKDIR /go/qyyh-go

RUN go build main.go

FROM docker.io/jeanblanchard/alpine-glibc as main

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn \
    Mysqlhost=host:3306 \
    Mysqluser=root \
    Mysqlpwd=Zxc940429+++ \
    Mysqldbname=qyyh

RUN apk --update add tzdata && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone

RUN date

WORKDIR /go

COPY --from=builder /go/qyyh-go/main /go
COPY --from=builder /go/qyyh-go/assets /go/assets

CMD ["/go/main"]
