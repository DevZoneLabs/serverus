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

# Stage 2: Use Alpine with Google Chrome
FROM alpine:latest

# Install Google Chrome and necessary dependencies
RUN apk add --no-cache \
    chromium \
    chromium-chromedriver \
    nss \
    freetype \
    harfbuzz \
    ca-certificates \
    ttf-freefont \
    libx11 \
    libxrender \
    libxrandr \
    libxi \
    mesa-gl

# Set the path for Chrome binary
ENV CHROME_BIN=/usr/bin/chromium-browser

# Create the application directory in the new image
RUN mkdir /app

# Copy the compiled Go binary from the builder stage
COPY --from=builder /app/serverusBotServer /app/

# Set the command to run your Go application
ENTRYPOINT ["/app/serverusBotServer"]
