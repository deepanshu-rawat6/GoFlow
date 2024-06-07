# Stage 1: Build stage
FROM golang:1.22-alpine AS BuildStage

# Set the working directory
WORKDIR /app

# Copy and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

EXPOSE 8080

# Build the Go application
RUN go build -o /build cmd/api/main.go

# Stage 2: Final stage
FROM alpine:latest

# Set the working directory
WORKDIR /

# Copy the binary from the build stage
COPY --from=BuildStage /build /build

EXPOSE 8080

# Set the entrypoint command
ENTRYPOINT ["/build"]