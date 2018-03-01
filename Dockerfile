FROM golang:alpine

RUN apk --no-cache add git bash wget curl
RUN go get github.com/codegangsta/gin

RUN go get -u github.com/gorilla/mux
RUN go get -u github.com/lib/pq
RUN go get -u github.com/jinzhu/gorm

WORKDIR /go/src/riggedstars

CMD gin run main.go

EXPOSE 3000
