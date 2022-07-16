echo:
	docker-compose exec echo sh
db:
	docker-compose exec db mysql --user=user --password=secret