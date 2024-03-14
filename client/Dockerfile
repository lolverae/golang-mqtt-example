FROM golang:1.22.0-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o main

FROM alpine:latest

COPY --from=builder /app/main /app/main

WORKDIR /app

EXPOSE 1883

ENTRYPOINT ["./main"]

