include .env

db_status:
	goose -dir=./migrations postgres "host=${DB_HOST} port=${DB_PORT} user=${DB_USER} password=${DB_PASSWORD} dbname=${DB_NAME}" status

migrate_up:
	goose -dir=./migrations postgres "host=${DB_HOST} port=${DB_PORT} user=${DB_USER} password=${DB_PASSWORD} dbname=${DB_NAME}" up

migrate_down:
	goose -dir=./migrations postgres "host=${DB_HOST} port=${DB_PORT} user=${DB_USER} password=${DB_PASSWORD} dbname=${DB_NAME}" down
