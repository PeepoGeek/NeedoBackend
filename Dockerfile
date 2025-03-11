FROM golang:1.20-alpine

WORKDIR /app

# Copiamos los archivos de go para descargar dependencias
COPY go.mod .
RUN go mod download

# Copiamos el resto del código
COPY . .

# Compilamos la aplicación
RUN go build -o server ./cmd/server

# Exponemos el puerto
EXPOSE 8080

# Comando de ejecución
CMD ["/app/server"] 