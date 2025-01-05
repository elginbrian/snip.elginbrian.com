# Step 1: Build the Go binary
FROM golang:1.20-alpine as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Go Modules manifests
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod tidy

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o main .

# Step 2: Create a smaller image for running the app
FROM alpine:latest  

# Install necessary dependencies (e.g., for PostgreSQL client)
RUN apk --no-cache add ca-certificates

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the Pre-built binary from the builder stage
COPY --from=builder /app/main .

# Expose port (must match the one in the config)
EXPOSE 8080

# Command to run the executable
CMD ["./main"]