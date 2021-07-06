FROM golang:1.16.3

WORKDIR /go/src
COPY ./src /go/src

RUN go get github.com/labstack/echo/

CMD ["go", "run", "server.go"]