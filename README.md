# echo-docker

## setup
```
git clone git@github.com:You-saku/echo-docker.git
cd echo-docker
cp server/.env.example server/.env
docker-compose up -d --build

docker-compose exec echo sh
go run main.go
```

## Doc
https://echo.labstack.com/

## Todo
* make CRUD
* imple Auth 