# Build container
FROM golang:1.22 AS build-env
WORKDIR /app

## Download Go modules
COPY go.mod go.sum ./
RUN go mod download && go mod verify

## Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux \
  go build -v -o api .

# Test the application
FROM build-env AS test-env
RUN go test -v ./...

# Release container
FROM alpine:latest AS release-env
WORKDIR /

## Copy the binary file
COPY --from=build-env /app/api app/api

## Add curl for healthcheck
RUN apk --no-cache add curl

## Expose used ports
EXPOSE 8080 8443

ENV LOG Verbose

## Start the application on boot up
CMD [ "/app/api" ]

HEALTHCHECK --timeout=5s \
  CMD curl -f http://localhost:8080/ || exit 1