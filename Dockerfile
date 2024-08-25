# Stage 1: Build the Go application
FROM golang:1.21-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules manifest and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

RUN go mod vendor

# Build the Go application
RUN go build -o /monigo-app

# Stage 2: Create the runtime container
FROM alpine:latest

# Install necessary dependencies including Graphviz
RUN apk update && \
    apk --no-cache add \
    ca-certificates \
    graphviz \
    && apk add --no-cache \
    bash  # Adding bash might be useful for debugging purposes

# Copy the compiled binary from the builder stage
COPY --from=builder /monigo-app /monigo-app

# Expose the ports your application will run on
EXPOSE 8000 8080

# Command to run the application
ENTRYPOINT ["/monigo-app"]
