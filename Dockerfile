FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/game ./cmd/game

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/game /app/game
COPY --from=builder /app/assets /app/assets

EXPOSE 8080

CMD ["/app/game"] 