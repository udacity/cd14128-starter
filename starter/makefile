# Makefile

.PHONY: test test-unit test-integration

# Run all tests
test: test-unit test-integration

# Run unit tests
test-unit:
	go test ./... -v -short

# Run integration tests
test-integration:
	go test ./tests/integration/... -v

# Create test database
create-test-db:
	PGPASSWORD=postgres psql -h localhost -U postgres -c "DROP DATABASE IF EXISTS cms_test;"
	PGPASSWORD=postgres psql -h localhost -U postgres -c "CREATE DATABASE cms_test;"

# Run integration tests with database setup
test-integration-full: create-test-db test-integration