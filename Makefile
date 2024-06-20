ifeq ($(POSTGRES_SETUP_TEST),)
	POSTGRES_SETUP_TEST := user=test password=test dbname=test host=localhost port=5432 sslmode=disable
endif

INTERNAL_PKG_PATH=$(CURDIR)/internal/pkg
MIGRATION_FOLDER=$(INTERNAL_PKG_PATH)/db/migrations

.PHONY: migration-create
# migration-create target creates a new migration file using goose
migration-create:
	goose -dir "$(MIGRATION_FOLDER)" create "$(name)" sql

.PHONY: test-migration-up
# test-migration-up target applies all pending migrations to the test database
test-migration-up:
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP_TEST)" up

.PHONY: test-migration-down
# test-migration-down target rolls back the last applied migration in the test database
test-migration-down:
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP_TEST)" down
