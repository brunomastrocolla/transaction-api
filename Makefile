build:
	go mod download
	go build -o ./transaction-api ./cmd/transaction-api/*.go

run:
	go run cmd/transaction-api/*.go server

migrate:
	go run cmd/transaction-api/*.go migrate

test:
	go test -covermode=count -coverprofile=count.out -v ./...

mock:
	@mkdir -p mocks
	mockgen -source=repository/repository.go -destination=mocks/mock_repository.go -package=mocks
	mockgen -source=service/service.go -destination=mocks/mock_service.go -package=mocks
	mockgen -source=handler/handler.go -destination=mocks/mock_handler.go -package=mocks

lint:
	./script/lint

.PHONY: build run migrate test mock lint
