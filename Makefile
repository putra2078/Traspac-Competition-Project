.PHONY: migrate-up migrate-down migrate-force

MIGRATION_PATH=migrations
DB_URL=postgresql://postgres:123456@localhost:5432/traspac_db?sslmode=disable

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

# Membuat folder domain baru dengan file entity, handler, repository, dan usecase
# Command: make domain name=users
domain:
	@if not defined name ( \
		echo ❌ Tolong isi nama domain, contoh: make domain name=users \
	) else ( \
		mkdir internal\domain\$(name) 2>nul && \
		(for %%f in (entity handler repository usecase) do ( \
			echo package $(name)>internal\domain\$(name)\%%f.go \
		)) && \
		echo ✅ Domain '$(name)' berhasil dibuat di internal\domain\$(name) \
	)

seed:
	psql "$(DB_URL)" -f migrations/seeders/seed_departments.sql
	psql "$(DB_URL)" -f migrations/seeders/seed_positions.sql

# Menjalankan aplikasi
run:
	go run cmd/server/main.go