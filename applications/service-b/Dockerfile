# Basic Dockerfile for Golang application
ARG GO_VERSION=1.21.5

FROM golang:${GO_VERSION}-alpine AS builder

WORKDIR /app

COPY . .

RUN go build -o /app/app main.go

FROM scratch

COPY --from=builder /app/app /app/app

ENTRYPOINT ["/app/app"]
