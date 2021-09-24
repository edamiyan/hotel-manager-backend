# hotel-manager

# Запуск
## Задайте ваш адрес и порт в файле config/confgis.yml
```
port: "8080"
URL: "0.0.0.0"

db:
  Host: "localhost"
  Port: "5432"
  Username: "postgres"
  DBName: "postgres"
  SSLMode: "disable"

```
## Установите и запустите БД и миграции
```
docker pull postgres
docker run --name=<db_name> -e POSTGRES_PASSWORD='<your_password>' -p 5432:5432 -d postgres
migrate create -ext sql -dir ./schema -seq init
migrate -path ./schema -database 'postgres://postgres:<your_password>@localhost:5432/postgres?sslmode=disable' up
migrate -path ./schema -database 'postgres://postgres:<your_password>@localhost:5432/postgres?sslmode=disable' down
```
## Скомпилируйте и запустите
```
go build cmd/main.go
./main
```
