FROM golang:1.24-alpine3.21 AS builder

WORKDIR /build

COPY ./cmd ./cmd

COPY ./vendor ./vendor
COPY ./go.mod ./go.sum ./

COPY ./internal ./internal

RUN go build -o app ./cmd/main.go

FROM alpine:3.21

WORKDIR /myapp

COPY --from=builder /build/app ./

EXPOSE 6080

CMD [ "/myapp/app" ]
