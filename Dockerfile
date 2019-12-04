FROM golang:1.13.0-alpine

EXPOSE 3001

RUN mkdir -p /go/src/app
COPY . /go/src/app
WORKDIR /go/src/app

RUN apk update && \
    apk add --no-cache git && \
    go get -u github.com/codegangsta/gin
CMD gin -i run
