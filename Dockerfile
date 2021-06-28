# Builder Stage
FROM golang:1.16.4-alpine3.13 as builder

# Copy application to the container
COPY . /go/src/users-srv

# Get dependencies
WORKDIR /go/src/users-srv
RUN go get -d -v ./...

# Build the application
RUN mkdir -p /go/src/users-srv/build
WORKDIR /go/src/users-srv/build

RUN go build -o ./users-service ../cmd/users/main.go

# Run Stage
FROM alpine:3.9

COPY --from=builder /go/src/users-srv/build /app/users-srv

ENV SERVER_PORT="8000"
EXPOSE 8000

# Run application
WORKDIR /app/users-srv
CMD ["./users-service"]