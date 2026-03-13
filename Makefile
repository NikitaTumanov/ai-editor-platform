goose-install:
	go install github.com/pressly/goose/v3/cmd/goose@latest

goose-create-migration:
	goose create new_migration sql

goose-up:
	goose up

goose-down:
	goose down

goose-status:
	goose status