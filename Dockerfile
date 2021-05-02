FROM golang:1.16.3-alpine as build

WORKDIR /go/app

COPY src .

RUN go build -o app

FROM alpine

WORKDIR /app

COPY --from=build /go/app/app .

CMD ["./app"]
