# Dockerfile
FROM golang:latest

# Install necessary tools
RUN go install github.com/air-verse/air@latest
RUN go install github.com/go-delve/delve/cmd/dlv@latest

EXPOSE 8080 2345

# Set the Current Working Directory inside the container
WORKDIR /app

RUN go mod init app && go mod tidy
# Copy go.mod and go.sum files
COPY app .

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the rest of the application code
COPY app .

RUN mkdir logs

# Command to run the Air live reload tool
CMD ["air", "-c", ".test.toml"]
