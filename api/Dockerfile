FROM golang:1.20-bullseye as builder

# Create and change to the app directory.
WORKDIR /app

# Download necessary Go modules
COPY go.mod /app
COPY go.sum /app
RUN go mod download

COPY . /app

WORKDIR /app/cmd

# Build the binary.
RUN go build -v -o server

FROM debian:buster-slim
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*

# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/cmd/server /app/server

EXPOSE 5000

# Run the web service on container startup.
CMD ["/app/server"]
