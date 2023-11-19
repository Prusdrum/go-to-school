GO_VERSION=1.21.4

install:
	go mod tidy

run-cli:
	go run ./cmd/cli/main.go

run-rest:
	go run ./cmd/rest-api/main.go

migrate-new:
	migrate create -ext sqlite3 -dir database/migration/ -seq name_here

migrate-up:
	go run ./cmd/db/main.go
	# migrate -database ./database/go-to-school-db.db -path ./database/migration/ up