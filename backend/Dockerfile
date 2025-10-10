# Use official Golang image as the base
FROM golang:1.24.7-alpine

WORKDIR /app

# Install git (if needed for go mod)
RUN apk add --no-cache git

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the backend source code
COPY . .

# Build the backend server
RUN go build -o server ./cmd/server

# Expose the backend port
EXPOSE 8080

# Run the backend server
CMD ["/app/server"]
