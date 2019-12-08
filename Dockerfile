FROM golang:1.13.0-alpine

EXPOSE 3001

RUN mkdir -p /go/src/app
COPY . /go/src/app
WORKDIR /go/src/app

RUN apk update && \
    apk add --no-cache git && \
    go get -u github.com/codegangsta/gin && \
    go get -u github.com/jinzhu/gorm && \
    go get -u github.com/jinzhu/gorm/dialects/postgres
CMD gin -i run
