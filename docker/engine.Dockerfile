FROM golang:1.21.3-bullseye AS build

# this is a multi stage build to make images lighter. 
# to make images from dockerfile:
# docker build -f docker/engine.Dockerfile .
# docker run -i -t c6dfde099801

# More lightweight
RUN apt-get update && apt-get install build-essential gcc -y

# Need to download librdkafka
ENV CGO_ENABLED 1

# Set up temporary build directory in container
WORKDIR /tmp/engine
COPY ./engine .
RUN go mod download

# Build Go 
RUN go build -o engine

# Start fresh from a smaller image
# Might need to change bullseye-slim to bullseye. idk though
FROM bitnami/minideb:bullseye

COPY ./kafka-topic-list.conf /kafka-topic-list.conf
COPY --from=build /tmp/engine/engine /engine

# Run the binary program
# the CMD will be run in docker-compose
# CMD ["/engine"]

# OLD CODE FOR ALPINE JUST HERE FOR REFERENCE

# FROM golang:1.21.3-alpine3.18 AS build

# # this is a multi stage build to make images lighter. 
# # to make images from dockerfile:
# # docker build -f docker/engine.Dockerfile .
# # docker run -i -t c6dfde099801

# # More lightweight
# RUN apk add --no-cache gcc musl-dev

# # Need to download librdkafka
# ENV CGO_ENABLED 1

# # Set up temporary build directory in container
# WORKDIR /tmp/engine
# COPY ./engine .
# RUN go mod download

# # Build Go (-tags musl is for alpine linux build)
# RUN go build -tags musl -ldflags '-extldflags "-static"' -o engine

# # Start fresh from a smaller image
# FROM alpine:3.18

# COPY ./kafka-topic-list.conf /kafka-topic-list.conf
# COPY --from=build /tmp/engine/engine /engine

# # Run the binary program
# # the CMD will be run in docker-compose
# CMD ["/engine"]