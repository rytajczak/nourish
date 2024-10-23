.PHONY: all
all: build up

.PHONY: build
build:
	docker-compose build

.PHONY: up
up:
	docker-compose up -d

.PHONY: down
down:
	docker-compose down

.PHONY: clean
clean:
	docker system prune --volumes -f
	docker-compose down --volumes

.PHONY: reset
reset: down clean