# Stage 1: Build the Go application
FROM golang:1.22 AS builder

# Create the application directory
RUN mkdir /app

# Copy the Go application source code to the container
COPY . /app

# Set the working directory
WORKDIR /app

# Build the Go application
RUN CGO_ENABLED=0 go build -o serverusBotServer ./cmd/

# Ensure the binary is executable
RUN chmod +x /app/serverusBotServer


# Stage 2: Use chromedp/headless-shell for a minimal headless Chromium environment
FROM chromedp/headless-shell:latest

# Set the path for the Chrome binary (chromedp/headless-shell already has this set up)
ENV CHROME_BIN=/headless-shell/headless-shell

# Create the application directory in the new image
RUN mkdir /app

# Copy the compiled Go binary from the builder stage
COPY --from=builder /app/serverusBotServer /app/

# Set the command to run your Go application
CMD ["/app/serverusBotServer"]
