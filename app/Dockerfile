FROM golang:alpine AS builder

WORKDIR /app
COPY . .
RUN go build -o app

FROM alpine:latest

RUN adduser -D appuser
USER appuser

WORKDIR /home/appuser
COPY --from=builder /app/app .

EXPOSE 8080
CMD ["./app"]