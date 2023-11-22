include .env

# Сборка исполняемого файла сервера
build:
	if [ -f "${BINARY}" ]; then \
		rm ${BINARY}; \
		echo "Deleted ${BINARY}"; \
	fi
	@echo "Building binary..."
	go build -o ${BINARY} cmd/server/*.go

# Запуск серверного приложения
run: build
	./${BINARY}

# Поднятие контейнера базы данных
container_up:
	docker run --name ${DB_DOCKER_CONTAINER} -p ${DB_PORT}:${DB_PORT} -e POSTGRES_USER=${USER} -e POSTGRES_PASSWORD=${PASSWORD} -d postgres:12-alpine

# Создание базы данных
create_db:
	docker exec -ti ${DB_DOCKER_CONTAINER} createdb --username=${USER} --owner=${USER} ${DB_NAME}

# Запуск контейнера базы данных
container_start:
	docker start ${DB_DOCKER_CONTAINER}

# Остановка контейнеров
container_stop:
	if [ $$(docker ps -q) ]; then \
		echo "stopping..."; \
		docker stop $$(docker ps -q); \
	else \
		echo "no containers running"; \
	fi