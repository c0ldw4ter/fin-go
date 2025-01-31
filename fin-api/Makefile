.PHONY: run migrate

run:
	  docker-compose up --build

migrate:
	  docker-compose exec app goose -dir migrations postgres "host=db port=5432 user=${DB_USER} dbname=${DB_NAME} sslmode=disable password=${DB_PASSWORD}" up