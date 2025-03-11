FROM golang:1.20-alpine

WORKDIR /app

# Install swag CLI
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Copiamos los archivos de go para descargar dependencias
COPY go.mod .
RUN go mod download

# Copiamos el resto del código
COPY . .

# Generamos la documentación Swagger
RUN swag init -g cmd/server/main.go

# Compilamos la aplicación
RUN go build -o server ./cmd/server

# Exponemos el puerto
EXPOSE 8080

# Comando de ejecución
CMD ["/app/server"] 