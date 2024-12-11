# Use a lightweight Go image as the base
FROM golang:1.22-alpine AS builder

# Set the working directory
WORKDIR /github.com/ANISH4623/

# Copy the Go source code
COPY . .

# Build the Go application
RUN go build -o main .

# Use a smaller image for the final image
FROM alpine:latest

# Copy the built binary from the builder stage
COPY --from=builder /github.com/ANISH4623/main /github.com/ANISH4623/main

# Set the working directory
WORKDIR /github.com/ANISH4623/

# Expose the port on which the Gin app will listen
EXPOSE 8080

# Command to run the Gin app
CMD ["./main"]
