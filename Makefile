setup:
	docker-compose up -d --build
up:
	docker-compose up -d
down:
	docker-compose down
ps:
	docker-compose ps
destroy:
	docker-compose down -v
start:
	docker-compose exec go go run main.go
go:
	docker-compose exec go sh
db:
	docker-compose exec db mysql --user=user --password=secret
create-migration: # nameを複数形にすること
	migrate create -ext sql -dir database/migrations -seq create_$(name)_table
test:
	docker-compose exec go go test ./test -v
