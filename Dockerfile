
# syntax=docker/dockerfile:1

# Use minimal base image
FROM alpine:latest

# Install bash in the final image
RUN apk update && apk add --no-cache bash

# Add a user for security
RUN adduser -D appuser

# Set working dir
WORKDIR /app

# Argument for the challenge binary
ARG CHALLENGE_BIN

# Copy the specific challenge binary into the image
COPY ${CHALLENGE_BIN} ./linux-learner

# Use non-root user
USER appuser

# Run the binary
ENTRYPOINT ["./linux-learner"]
