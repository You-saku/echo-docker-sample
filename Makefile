setup:
	docker-compose up -d --build
up:
	docker-compose up -d
down:
	docker-compose down
break:
	docker-compose down -v
start:
	docker-compose exec echo go run main.go
go:
	docker-compose exec echo sh
db:
	docker-compose exec db mysql --user=user --password=secret