
# Use an official Golang runtime as a parent image
FROM docker.io/library/golang:latest

# Set the working directory to /go/src/app
WORKDIR /go/src/app

# Copy the current directory contents into the container at /go/src/app
COPY . /go/src/app

# Install required libraries
RUN go get -u github.com/go-sql-driver/mysql
RUN go get github.com/gin-gonic/gin
RUN go get "github.com/gin-contrib/cors"

ENV DB_HOST=172.17.0.1 \
    DB_PORT=3306 \
    SERVER_PORT=8080

# Expose port 8080 for the application
EXPOSE ${SERVER_PORT}

# Build the Go application
RUN go build -o main .

# Run the Go application when the container starts with the command go run .
CMD ["./main"]
