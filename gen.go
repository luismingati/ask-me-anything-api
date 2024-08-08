package gen

//go:generate sqlc generate -f ./internal/store/pgstore/sqlc.yml
//go:generate go run ./cmd/tools/terndotenv/main.go
