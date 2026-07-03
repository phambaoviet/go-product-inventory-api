include .env
export
MIGRATION_DIRS = internal/db/migrations
CONN_STRING = postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)


import_db:
	docker exec -i postgres-db psql -U root -d master_golang < ./backupdb-master-golang.sql
export_db:
	docker exec -i postgres-db pg_dump -U root -d master_golang > ./backupdb-master-golang.sql
server:
	go run .
migrate-create:
	migrate create -ext sql -dir $(MIGRATION_DIRS) -seq $(name)
migrate-up:
	migrate -path $(MIGRATION_DIRS) -database $(CONN_STRING) up
migrate-down:
	migrate -path $(MIGRATION_DIRS) -database $(CONN_STRING) down 1
migrate-down-n:
	migrate -path $(MIGRATION_DIRS) -database $(CONN_STRING) down $(n)
migrate-force:
	migrate -path $(MIGRATION_DIRS) -database $(CONN_STRING) force $(version)
migrate-drop:
	migrate -path $(MIGRATION_DIRS) -database $(CONN_STRING) drop
migrate-goto:
	migrate -path $(MIGRATION_DIRS) -database $(CONN_STRING) goto $(version)
.PHONY: import_db export_db server migrate-create migrate-up migrate-down migrate-force migrate-drop migrate-goto migrate-down-n

