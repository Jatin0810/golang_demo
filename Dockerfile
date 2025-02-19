
# # # Run
# # CMD ["/golang_mongodb"]

# # Use the official Go image as a base
# FROM golang:1.23

# # Set the working directory inside the container
# WORKDIR /app

# # Copy go.mod and go.sum first to leverage Docker layer caching
# COPY go.mod go.sum ./

# # Download Go dependencies
# RUN go mod download

# # Copy the entire project into the container
# COPY . .

# # Set the working directory to the correct directory containing main.go
# WORKDIR /app/cmd/golang_mongodb

# # Build the Go application
# RUN CGO_ENABLED=0 GOOS=linux go build -o /golang-mongodb-docker

# # Expose the required port
# EXPOSE 9090

# # Copy the built binary from the builder stage
# # COPY --from=builder /golang-mongodb-docker .

# # Copy the config file
# COPY configs/configs.yaml /configs/configs.yaml

# # Run the compiled binary
# CMD ["/golang-mongodb-docker",  "-config", "/configs/configs.yaml" ]


# Telling to use Docker's golang ready image
FROM golang
# Name and Email of the author 
MAINTAINER Osama Elmashad <elmashad285@gmail.com>
# Create app folder 
RUN mkdir /app
# Copy our file in the host contianer to our contianer
ADD . /app
# Set /app to the go folder as workdir
WORKDIR /app
# Generate binary file from our /app
RUN go build
# Expose the port 3000
EXPOSE 3000:3000
# Run the app binarry file 
CMD ["./app", "-config", "/configs/configs.yaml"]