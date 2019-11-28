FROM golang:1.12.0-alpine

EXPOSE 3001

ENV GOPATH /go
ENV PATH $PATH:$GOPATH/bin

RUN mkdir -p /go/src/app
COPY . /go/src/app
WORKDIR /go/src/app

RUN apk update && \
    apk add --no-cache git && \
    go get -u github.com/codegangsta/gin && \
    go get -u github.com/golang/dep/cmd/dep && \
    dep init || dep ensure
CMD gin -i run
