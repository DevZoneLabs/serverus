# Stage 1: Build the Go application
FROM golang:1.22 as builder

WORKDIR /app

# Copy all project files
COPY . .

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -o serverusBotServer ./cmd/

# Stage 2: Create a minimal image with Alpine and Google Chrome
FROM alpine:latest

# Install required dependencies and download Chrome
RUN apk add --no-cache \
    wget \
    nss \
    freetype \
    harfbuzz \
    ca-certificates \
    ttf-freefont \
    libx11 \
    libxrender \
    libxrandr \
    libxi \
    mesa-gl \
    dumb-init \
    fontconfig \
    libc6-compat

# Download and install Google Chrome
RUN wget https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb && \
apk add --no-cache --virtual .build-deps dpkg && \
dpkg -i google-chrome-stable_current_amd64.deb || apk add -f && \
rm google-chrome-stable_current_amd64.deb && \
apk del .build-deps

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
