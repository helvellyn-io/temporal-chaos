# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:latest

# Add Maintainer Info
LABEL maintainer="Dylan Johnson <dylan@helvellyn.io>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy everything from the current directory to the Working Directory inside the container
COPY . .

WORKDIR /app/cmd
# Build the Go app
RUN go build -o main .

RUN cp main ../

WORKDIR /app

# Command to run the executable
CMD ["./main"]