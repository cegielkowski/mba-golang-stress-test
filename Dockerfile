# Etapa de construção
FROM golang:1.18-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o load-tester main.go

# Etapa de execução
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/load-tester .
ENTRYPOINT ["./load-tester"]
