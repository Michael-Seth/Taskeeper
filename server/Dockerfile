# Stage 1: Build Stage
FROM golang:1.23-alpine AS build

# Set the Current Working Directory inside the container
WORKDIR /go/app

# Copy go.mod and go.sum files from the server folder and download dependencies
COPY ./server/go.mod ./server/go.sum ./

# Download Go module dependencies
RUN go mod download

# Copy the entire project source (from the root)
COPY ./server /go/app/

# Build the Go application (assuming your main app is in the cmd folder)
# RUN go build -o /go/app/cmd/app ./cmd
RUN go build -mod=mod -o /go/app/cmd/app ./cmd

# Stage 2: Production Stage
FROM alpine:3.18

# Set a working directory
WORKDIR /app

# Copy the compiled binary from the build stage
COPY --from=build /go/app/cmd/app .

# Expose application port
EXPOSE 8080

# Command to run the executable
CMD ["./app"]
