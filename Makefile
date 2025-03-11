# Compila la aplicación dentro del contenedor y construye la imagen
build:
	docker build -t myproject:latest .

# Levanta todos los servicios (app y db)
run:
	docker-compose up -d

# Detiene todos los servicios
stop:
	docker-compose down

# Limpia la imagen construida (opcional)
clean:
	docker rmi myproject:latest || true 