FROM golang:1.22 as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o serverusBotServer ./cmd/

RUN chmod +x /app/serverusBotServer

# Build a tiny docker image with Chrome
FROM alpine:latest

# Install Chrome dependencies
RUN apk add --no-cache \
    chromium \
    nss \
    freetype \
    harfbuzz \
    ttf-freefont

# Set the path for Chrome binary
ENV CHROME_BIN=/usr/bin/chromium-browser

RUN mkdir /app

COPY --from=builder /app/serverusBotServer /app

CMD ["/app/serverusBotServer"]