# Stage 1: Build your Go application
FROM golang:1.22 as builder

# Set the working directory
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of your application code
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 go build -o serverusBotServer ./cmd/

FROM docker.io/chromedp/headless-shell:latest

# Install the dependencies
RUN apt-get update && \
    apt-get install -y ca-certificates dumb-init unzip zip wget zlib1g-dev openssh-client && \
    rm -rf /var/lib/apt/lists/*

# Create the application directory in the new image
RUN mkdir /app

# Copy the compiled Go binary from the builder stage
COPY --from=builder /app/serverusBotServer /app/

ENTRYPOINT ["dumb-init", "--"]

CMD ["/app/serverusBotServer"]
