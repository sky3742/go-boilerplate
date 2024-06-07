migrate:
	go run cmd/db/migrate/migrate.go

rollback:
	go run cmd/db/rollback/rollback.go

start:
	go run cmd/api/main.go

watch:
	air
