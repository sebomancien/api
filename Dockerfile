# Build container
FROM golang:1.22 AS build-env
WORKDIR /app

## Download Go modules
COPY go.mod go.sum ./
RUN go mod download && go mod verify

## Copy the source code
COPY *.go ./

# Build the application
RUN CGO_ENABLED=0 GOOS=linux \
  go build -o api -v ./...

# Test the application
FROM build-env AS test-env
RUN go test -v ./...

# Release container
FROM alpine:latest AS release-env
WORKDIR /

## Copy the binary file
COPY --from=build-env /app/api app/api

## Expose used ports
EXPOSE 8080 8443

ENV LOG Verbose

## Start the application on boot up
CMD [ "/app/api" ]