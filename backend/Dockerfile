FROM golang:1.16.3

RUN go get github.com/labstack/echo/...

WORKDIR /app
ADD . /app

CMD ["go", "run", "main.go"]