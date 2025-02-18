# # A sample microservice in Go packaged into a container image.
# FROM golang:1.20

# # Set destination for COPY
# WORKDIR /app/cmd/golang_mongodb

# # Download Go modules
# COPY go.mod go.sum ./
# RUN go mod download 

# # Copy the source code. Note the slash at the end, as explained in
# # https://docs.docker.com/engine/reference/builder/#copy
# COPY *.go ./

# # Build
# RUN CGO_ENABLED=0 GOOS=linux go build -o /golang_mongodb/cmd/golang_mongodb

# # Optional:
# # To bind to a TCP port, runtime parameters must be supplied to the docker command.
# # But we can document in the Dockerfile what ports
# # the application is going to listen on by default.
# # https://docs.docker.com/engine/reference/builder/#expose
# EXPOSE 9090

# # Run
# CMD ["/golang_mongodb"]

# Use the official Go image as a base
FROM golang:1.23

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum first to leverage Docker layer caching
COPY go.mod go.sum ./

# Download Go dependencies
RUN go mod download

# Copy the entire project into the container
COPY . .

# Set the working directory to the correct directory containing main.go
WORKDIR /app/cmd/golang_mongodb

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o /golang-mongodb-docker

# Expose the required port
EXPOSE 9090

# Run the compiled binary
CMD ["/golang-mongodb-docker"]
