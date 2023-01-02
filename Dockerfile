# syntax=docker/dockerfile1

# Build Stage

FROM golang:1.18-alpine AS build

WORKDIR /app

COPY go.mod .

RUN go mod download

COPY . /app/

# RUN apk add --no-cache gcc musl-dev

# RUN go test ./...

RUN go build -o /main

CMD ["/main"]