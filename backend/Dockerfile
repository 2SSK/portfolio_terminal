# Use a small Go image
FROM golang:1.24-alpine AS builder

# Set the working directory
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./
RUN go mod download

# Copy the application source code
COPY . .

# Build the application
RUN go build -o server .

# Use a minimal base image for deployment
FROM alpine:latest

# Set the working directory
WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/server .

# Expose the application port
EXPOSE 3000

# Run the application
CMD ["./server"]
