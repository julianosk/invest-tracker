# Start from the official Golang image
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the rest of the source code
COPY . .

# Build the Go app
RUN go build -o invest-tracker main.go

# Use a minimal image for running
FROM alpine:latest

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/invest-tracker .

# Expose the port
EXPOSE 8080

# Set environment variables (can be overridden by docker-compose)
ENV PORT=8080

# Run the application
CMD ["./invest-tracker"]