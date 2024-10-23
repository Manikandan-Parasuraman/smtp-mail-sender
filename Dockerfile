# Step 1: Use the official Golang image as a builder
FROM golang:1.22.5 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum for dependency management
COPY go.mod go.sum ./

# Download dependencies
RUN go mod tidy

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN go build -o smtp-mail-sender .

# Step 2: Create a minimal runtime image
FROM debian:buster-slim

# Set environment variables for SMTP
ENV SMTP_HOST=smtp.gmail.com
ENV SMTP_PORT=587
ENV SMTP_USER=your-email@gmail.com
ENV SMTP_PASS=your-email-password

# Set working directory in runtime image
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/smtp-mail-sender /app/

# Copy the HTML templates directory
COPY --from=builder /app/templates /app/templates

# Expose any ports the app may use (not necessary for SMTP but added for general use)
EXPOSE 8080

# Run the binary
CMD ["go run main.go"]
