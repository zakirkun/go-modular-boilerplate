# Step 1: Use the official Golang image to build the app
# This uses the latest Go version
FROM golang:1.23.1-alpine3.20 AS builder

# Step 2: Set the working directory inside the container
WORKDIR /app

# Step 3: Copy the go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Step 4: Download all Go module dependencies
# Dependencies will be cached if the go.mod and go.sum files haven't changed
RUN go mod download

# Step 5: Copy the rest of the application code to the container
COPY . .

# Step 6: Build the Go app
RUN go build -o main ./cmd/app/

# Step 7: Use a minimal base image for running the app
FROM alpine:latest

# Step 8: Set the working directory for the minimal base image
WORKDIR /app

# Step 9: Copy the Go app from the builder stage
COPY --from=builder /app/main .

# Step 10: Copy Configuration
COPY config.toml .

# Step 11: Command to run the app
CMD ["./main", "-c", "config.toml"]