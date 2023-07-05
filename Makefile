include .env
export $(shell sed 's/=.*//' .env)

DATABASE_URL := $(DB_TYPE)://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable

migrate-create:
	migrate create -ext sql -dir ./src/config/database/migrations -seq $(NAME)

migrate-up:
	migrate -path ./src/config/database/migrations -database "$(DATABASE_URL)" -verbose up

migrate-down:
	migrate -path ./src/config/database/migrations -database "$(DATABASE_URL)" -verbose down
