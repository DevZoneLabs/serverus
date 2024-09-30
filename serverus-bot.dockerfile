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

# Stage 2: Use chromedp/headless-shell for a minimal headless Chromium environment
FROM chromedp/headless-shell:latest

# Install CA certificates and curl
RUN apk --no-cache add ca-certificates curl && update-ca-certificates

# Test connectivity to Discord API
RUN curl -v https://discord.com/api/v9/gateway

# Create the application directory in the new image
RUN mkdir /app

# Copy the compiled Go binary from the builder stage
COPY --from=builder /app/serverusBotServer /app/

# Set the command to run your Go application
ENTRYPOINT ["/app/serverusBotServer"]
