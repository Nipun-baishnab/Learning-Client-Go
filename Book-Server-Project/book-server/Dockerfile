# Use an official Golang runtime as a parent image
FROM golang:1.20-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod ./
COPY go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o book-server .

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD ["./book-server"]
