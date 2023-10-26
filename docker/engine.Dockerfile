FROM golang:1.21.3-alpine3.18 AS build

# this is a multi stage build to make images lighter. 
# to make images from dockerfile:
# docker build -f docker/engine.Dockerfile .

# More lightweight
RUN apk add --no-cache git

# Set up temporary build directory in container
WORKDIR /tmp/engine
COPY ./engine .
RUN go mod download

# Build Go
RUN go build -o engine

# Start fresh from a smaller image
FROM alpine:3.9 

COPY --from=build /tmp/engine/engine /engine

# Run the binary program
# the CMD will be run in docker-compose
# CMD ["/engine"]