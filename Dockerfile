FROM golang:1.16.3

WORKDIR /go/app

RUN go get github.com/go-sql-driver/mysql

COPY src .

# RUN go build -o app

CMD ["go", "run", "main.go"]
