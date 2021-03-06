FROM golang:alpine

RUN apk --no-cache add git
RUN go get -u github.com/codegangsta/gin github.com/gorilla/mux github.com/lib/pq github.com/jinzhu/gorm github.com/gorilla/websocket golang.org/x/crypto/bcrypt github.com/dgrijalva/jwt-go github.com/gorilla/handlers

WORKDIR /go/src/riggedstars

CMD gin run main.go

EXPOSE 3000
