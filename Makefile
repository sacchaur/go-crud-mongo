all: install

install:
	# Stop Docker
	docker-compose down
	# Remove Docker images
	docker image rm go-crud-mongo-go-mongodb-api --force
	# Start Docker
	docker-compose up -d --force-recreate
	@echo "Application is up and running on http://localhost:3000"