setup:
	docker-compose up -d --build
	migrate -database "mysql://user:secret@tcp(127.0.0.1:3306)/develop?multiStatements=true" -path=database/migrations up 3
	docker-compose exec -T db mysql --user=user --password=secret develop < database/dev/setup.sql
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
create-table: # nameを複数形にすること
	migrate create -ext sql -dir database/migrations -seq create_$(name)_table
test:
	docker-compose exec go go test ./test -v
