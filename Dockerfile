# first stage build the app
FROM golang AS builder

# Add Maintainer Info
LABEL maintainer="Ariel Henryson <ariel.henryson@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the source from the current directory to the
# Working Directory inside the container
COPY . .

# Build the binary.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Start a new stage
FROM alpine:latest

# https://hackernoon.com/alpine-docker-image-with-secured-communication-ssl-tls-go-restful-api-128eb6b54f1f
RUN apk --no-cache add ca-certificates

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Copy the demo yaml file
COPY --from=builder /app/demo.yaml .

# start the engine app
CMD ["./main", "demo.yaml"]
