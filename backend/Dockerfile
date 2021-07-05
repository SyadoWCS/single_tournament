FROM golang:1.16.3

RUN go get github.com/labstack/echo/...

WORKDIR /backend
ADD . /backend

CMD ["go", "run", "main.go"]