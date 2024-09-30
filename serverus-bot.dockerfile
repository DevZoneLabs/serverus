# Stage 1: Build the Go application
FROM golang:1.22 as builder

WORKDIR /app

# Copy all project files
COPY . .

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -o serverusBotServer ./cmd/

# Stage 2: Create a Debian-based image with Google Chrome
FROM debian:bullseye-slim

# Install dependencies
RUN apt-get update && apt-get install -y \
    wget \
    gnupg2 \
    apt-transport-https \
    ca-certificates \
    curl \
    dumb-init \
    --no-install-recommends && \
    rm -rf /var/lib/apt/lists/*

# Download and install Google Chrome
RUN wget -q -O - https://dl.google.com/linux/linux_signing_key.pub | apt-key add - && \
    sh -c 'echo "deb [arch=amd64] http://dl.google.com/linux/chrome/deb/ stable main" > /etc/apt/sources.list.d/google-chrome.list' && \
    apt-get update && \
    apt-get install -y google-chrome-stable && \
    rm -rf /var/lib/apt/lists/*

# Add Google Chrome to PATH
ENV PATH="/usr/bin/google-chrome:${PATH}"

# Set environment variable for Chrome binary
ENV CHROME_BIN=/usr/bin/google-chrome

# Create /app directory and copy the Go binary from the builder stage
RUN mkdir /app
COPY --from=builder /app/serverusBotServer /app/

# Ensure the binary is executable
RUN chmod +x /app/serverusBotServer

# Use dumb-init as the entrypoint to handle signals and clean up processes properly
ENTRYPOINT ["/usr/bin/dumb-init", "--"]

# Command to run your Go server
CMD ["/app/serverusBotServer"]
