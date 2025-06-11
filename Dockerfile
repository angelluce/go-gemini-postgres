# Etapa de compilación
FROM golang:1.23.4-alpine AS builder

# Instalamos las dependencias
RUN apk add --no-cache gcc musl-dev

WORKDIR /app

# Copiamos solo los archivos necesarios para descargar dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copiamos el resto del código fuente
COPY . .

# Compilamos la aplicación
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Etapa final
FROM alpine:latest

# Instalamos los certificados CA para HTTPS
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copiamos solo el binario compilado desde la etapa anterior
COPY --from=builder /app/main .
COPY --from=builder /app/.env .

# Exponemos el puerto
EXPOSE 8080

# Ejecutamos la aplicación
CMD ["./main"]