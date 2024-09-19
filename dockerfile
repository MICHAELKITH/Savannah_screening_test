# Use the appropriate Go version image
FROM golang:1.23 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files first, then install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go app
RUN go build -o main .

# Expose port and define entrypoint
EXPOSE 8080
CMD ["./main"]



#docker build . -t go-containerized:latest
#docker image ls | grep go-containerized 