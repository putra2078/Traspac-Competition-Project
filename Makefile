.PHONY: migrate-up migrate-down migrate-force

MIGRATION_PATH=migrations
DB_URL=postgresql://postgres:123456@localhost:5432/hrm_db?sslmode=disable

# Menjalankan migrasi up
migrate-up:
	migrate -database "$(DB_URL)" -path $(MIGRATION_PATH) up

# Menjalankan migrasi down (rollback)
migrate-down:
	migrate -database "$(DB_URL)" -path $(MIGRATION_PATH) down

# Force migrasi ke versi tertentu jika terjadi error
migrate-force:
	migrate -database "$(DB_URL)" -path $(MIGRATION_PATH) force $(version)

# Membuat file migrasi baru
migrate-create:
	migrate create -ext sql -dir $(MIGRATION_PATH) -seq $(name)

# Menjalankan aplikasi
run:
	go run cmd/server/main.go