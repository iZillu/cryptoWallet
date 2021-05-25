build:
	docker-compose build crypto-wallet

run:
	docker-compose up crypto-wallet

migrate:
	migrate -path ./schema -database 'postgres://postgres:123123@localhost:5432/postgres?sslmode=disable' up
