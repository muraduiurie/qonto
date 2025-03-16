# Use the official Golang image as a build stage
FROM golang:1.23.5 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules files and download dependencies
COPY . .

# Download dependencies
RUN go mod download

# Change directory to the server
RUN cd ./cmd/server

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

# Use a minimal base image for the final container
FROM alpine:latest

# Set working directory inside the container
WORKDIR /root/

# Copy the built binary from the builder stage
COPY --from=builder /app/main .

# Run the application
CMD ["./main"]