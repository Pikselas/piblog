# Go development container
FROM mcr.microsoft.com/devcontainers/go:dev-1.22-bullseye

# Copy the local package files to the container's workspace.
COPY . /app

# Set the current directory inside the container
WORKDIR /app

# Build the Go app
RUN go build .

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the binary program produced by `go build`
ENTRYPOINT ["./main"]