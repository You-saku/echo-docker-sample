up:
	docker-compose up -d
down:
	docker-compose down
echo:
	docker-compose exec echo sh
db:
	docker-compose exec db mysql --user=user --password=secret