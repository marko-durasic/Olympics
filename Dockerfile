# Start with the official Golang image
FROM golang:1.22

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Copy the wait-for-it script
COPY wait-for-it.sh /wait-for-it.sh

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./wait-for-it.sh", "db:5432", "--", "./main"]
