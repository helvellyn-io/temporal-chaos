# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:latest

LABEL maintainer="Dylan Johnson <dylan@helvellyn.io>"

WORKDIR /app

COPY . .

WORKDIR /app/cmd

RUN go mod download

RUN go build -o main .

RUN cp main ../

WORKDIR /app

CMD ["./main"]