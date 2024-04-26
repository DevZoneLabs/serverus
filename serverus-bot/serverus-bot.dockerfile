FROM golang:1.22 as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o serverusBotServer ./cmd/

RUN chmod +x /app/serverusBotServer

# build a tiny docker image
FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/serverusBotServer /app

CMD [ "sh" ]