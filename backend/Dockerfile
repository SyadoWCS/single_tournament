FROM golang:1.16.3

WORKDIR /app
COPY . /app

RUN go get github.com/labstack/echo/v4
RUN go get github.com/labstack/echo/v4/middleware
RUN go get gorm.io/driver/mysql
RUN go get gorm.io/gorm
RUN go get golang.org/x/crypto/bcrypt@v0.0.0-20210322153248-0c34fe9e7dc2
RUN go get github.com/dgrijalva/jwt-go

CMD ["go", "run", "main.go"]