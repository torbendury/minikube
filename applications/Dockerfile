# Basic Dockerfile for Golang application
# This will be reused for "service-a" and "service-b" application
ARG GO_VERSION=1.21.5
ARG SERVICE_DIRECTORY=service-a

FROM golang:${GO_VERSION}-alpine AS builder

WORKDIR /app

COPY ${SERVICE_DIRECTORY}/ .

RUN go build -o /app/app main.go

FROM scratch

COPY --from=builder /app/app /app/app

CMD ["/app/app"]
