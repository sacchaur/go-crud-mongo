# Use the official Golang image as the base image
FROM golang:1.22.3-alpine

# Set the working directory inside the container
WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

# Copy the Go module files
COPY go.mod go.sum ./

# Download and install the Go dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN go build -o main .

# Set environment variables
# ENV PORT=$PORT
# ENV MONGODB_USERNAME=$MONGODB_USERNAME
# ENV MONGODB_PASSWORD=$MONGODB_PASSWORD
# ENV MONGODB_URI=$MONGODB_URI
# ENV MONGODB_DB_NAME=$MONGODB_DB_NAME
# ENV MONGODB_USER_COLLECTION=$MONGODB_USER_COLLECTION
# ENV MONGODB_TIMEOUT=$MONGODB_TIMEOUT

# Set the command to run the executable
CMD ["./main"]
