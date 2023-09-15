up:
	docker-compose up -d

down:
	docker-compose down

log:
	docker-compose logs -f

build:
	docker build -t buildgotodo --target deploy ./