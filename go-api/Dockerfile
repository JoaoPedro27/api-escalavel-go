# Estágio de build
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copia os arquivos de dependência
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copia o código e compila
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /api

# Estágio final (imagem leve)
FROM alpine:latest

WORKDIR /app

# Copia o binário do estágio anterior
COPY --from=builder /api /app/api

# Expõe a porta e executa
EXPOSE 8080
CMD ["/app/api"]