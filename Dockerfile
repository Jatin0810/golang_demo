FROM golang:latest

ENV Dev
ENV PORT=8002
WORKDIR /app
COPY go.mod /app
COPY go.sum /app

RUN go mod download

COPY . /app
ENTRYPOINT CompileDaemon --build="go build -o cmd/golang_mongodb/main" --command=./cmd/golang_mongodb/main