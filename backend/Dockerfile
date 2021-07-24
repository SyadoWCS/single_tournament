FROM golang:1.16.3

WORKDIR /app
COPY . /app

RUN go get github.com/labstack/echo
RUN go get gorm.io/driver/mysql
RUN go get gorm.io/gorm

CMD ["go", "run", "main.go"]