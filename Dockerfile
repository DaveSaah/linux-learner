# syntax=docker/dockerfile:1

# Use minimal Go base image for production
FROM golang:1.24.2-alpine AS builder

# Set working directory
WORKDIR /app

# Copy go modules and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go binary statically
RUN CGO_ENABLED=0 go build -o linux-learner ./cmd/linux-learner

# Final minimal image
FROM alpine:latest

# Install bash in the final image
RUN apk update && apk add --no-cache bash

# Add a user for security
RUN adduser -D appuser

# Set working dir and copy binary
WORKDIR /app
COPY --from=builder /app/linux-learner .

# Copy embedded YAML data if needed (you already use //go:embed)
# Not required if embedded correctly

# Use non-root user
USER appuser

# Run the binary
ENTRYPOINT ["./linux-learner"]
