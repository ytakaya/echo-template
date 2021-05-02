FROM golang:1.16.3 as builder

WORKDIR /go/app

COPY src/go.mod .
COPY src/go.sum .

RUN go mod download
COPY src .

RUN CGO_ENABLED=0 go build main.go

FROM alpine as production

WORKDIR /go/app

COPY --from=builder /go/app/main /go/app/main

EXPOSE 8080
CMD ["/go/app/main"]
