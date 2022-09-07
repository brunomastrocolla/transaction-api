# transaction-api

A REST-API server to manage accounts and financial transactions 

### Requirements (run/dev)
* Docker
* Docker Compose

### Requirements (dev)
* Go 1.17

## Docker Run
1. Building docker images
```bash
docker-compose -f docker-compose.yml build
```
2. Running
```bash
docker-compose -f docker-compose.yml up -d
```

## Integration Test
1. Create new account 
```bash
curl -i --request POST \
   --url http://localhost:8080/accounts \
   --header 'Content-Type: application/json' \
   --data '{ "document_number": "0123456789" }'
```
2. Find account by ID
```bash
curl -i --request GET --url http://localhost:8080/accounts/1
```
3. Create new transaction
```bash
curl -i --request POST \
  --url http://localhost:8080/transactions \
  --header 'Content-Type: application/json' \
  --data '{ "account_id": 1, "operation_type_id": 4, "amount": 123.45 }'
```

## Unit Test
1. Generate mocks
```bash
make mock
```
2. Running unit tests
```bash
make test
```

## Development
1. Setup postgres database
```bash
docker run --name=postgres -p 5432:5432 --env="POSTGRES_USER=user" --env="POSTGRES_PASSWORD=pass" --restart=unless-stopped --detach=true postgres:12-alpine
docker exec -it postgres psql -U user -c "create database transaction_api"
```
2. Set environment variables
```bash
cp example.env .env
vim .env
```
3. Build
```bash
make build
```
4. Run database migrate
```bash
make migrate
```
4. Run server
```bash
make run
```
4. Run linter
```bash
make lint
```

