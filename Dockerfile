# FROM golang:latest
FROM golang:alpine

WORKDIR /project

ENV PORT 5000
EXPOSE 5000

RUN apk add --update --no-cache make gcc g++ python git && \
    rm -rf /var/cache/apk/*

RUN go get github.com/tools/godep && \
    go install github.com/tools/godep
RUN go get github.com/gin-gonic/gin && \
    go install github.com/gin-gonic/gin
RUN go get github.com/jinzhu/gorm && \
    go get github.com/mattn/go-sqlite3 && \
    go get github.com/lib/pq

ADD . /project
RUN go install ./cmd/crd-controller
CMD ./bin/crd-controller
