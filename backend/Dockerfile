# Stage 1: Build the application
FROM golang:1.24-alpine AS builder

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project (excluding prisma/db/ via .dockerignore)
COPY . .

# Ensure prisma directory exists and copy schema
COPY prisma/schema.prisma ./prisma/

# Generate Prisma client for Linux
RUN GOOS=linux go run github.com/steebchen/prisma-client-go generate

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./main.go

# Stage 2: Create a lightweight runtime image
FROM alpine:latest

# Install ca-certificates for HTTPS requests
RUN apk --no-cache add ca-certificates

# Set working directory
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Copy Prisma schema (needed for runtime migrations, if applicable)
COPY --from=builder /app/prisma ./prisma/

# Expose the port
EXPOSE 3000

# Command to run the application
CMD ["./main"]
