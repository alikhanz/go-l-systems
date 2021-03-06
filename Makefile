build:
	go build -o bin/app cmd/main.go
run-go:
	go run cmd/main.go
test:
	go test ./...