FROM golang:latest
WORKDIR /project
ENV PORT 8080
EXPOSE 8080

RUN go get github.com/tools/godep && \
    go install github.com/tools/godep
RUN go get github.com/gin-gonic/gin && \
    go install github.com/gin-gonic/gin
RUN go get github.com/jinzhu/gorm && \
    go get github.com/mattn/go-sqlite3 && \
    go get github.com/lib/pq

ADD . /project
ENTRYPOINT  ["/usr/local/go/bin/go"]
CMD ["run", "./cmd/crd-controller/crud.go"]
