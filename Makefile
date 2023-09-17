up:
	docker-compose up -d

down:
	docker-compose down

log:
	docker-compose logs -f

build:
	docker build -t buildgotodo --target deploy ./
ssh:
	docker-compose exec -it app bash

migrate:
	mysqldef -u todo -p password -h todo-db -P 3306 todo < ./_tools/mysql/schema.sql

dry-migrate:
	mysqldef -u todo -p password -h todo-db -P 3306 todo --dry-run < ./_tools/mysql/schema.sql