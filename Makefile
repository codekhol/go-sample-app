.PHONY build:
build:
	docker compose build

.PHONY test:
test:
	docker compose up api-test

.PHONY run:
run:
	docker compose up --build
