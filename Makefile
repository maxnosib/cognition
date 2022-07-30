COVER_FILE ?= coverage.out
export PORT=:8080
export PGPASSWORD=postgres
export PG_HOST=localhost
export PG_PORT=54320
export PG_USER=postgres
export PG_PWD=postgres
export PG_DB=cognition

run:		## run application
	go run main.go

build: ## Build the project binary
	CGO_ENABLED=0 GOOS=linux go build -o main .

test: $(COVER_FILE) ## выводим покрытие тестами
	go test ./... -coverprofile=$(COVER_FILE)
	go tool cover -func=$(COVER_FILE) | grep ^total

cover_html: ## выводим в браузере покрытие тестами
	go tool cover -html=coverage.out

migration-init:		## init migration
	./migrago -c migration/config.yaml init

migration-up:		## roll up migrations
	./migrago -c migration/config.yaml up

infrastructure-up:
	sh ./scripts/start_infrastructure.sh

start-docker:		## run application
	docker-compose up -d

migration-init-local:		## init migration
	./migrago -c migration/config_local.yaml init

migration-up-local:		## roll up migrations
	./migrago -c migration/config_local.yaml up
