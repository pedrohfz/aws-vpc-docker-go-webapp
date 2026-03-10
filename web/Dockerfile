FROM golang:1.25-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o webserver .

FROM alpine:3.20
WORKDIR /app

COPY --from=builder /app/webserver /app/webserver
COPY --from=builder /app/templates /app/templates
COPY --from=builder /app/static /app/static

ENV BACKEND_URL=http://localhost:25000

EXPOSE 8080
CMD ["/app/webserver"]