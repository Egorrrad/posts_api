build:
	docker-compose build post-api

run:
	docker-compose up post-api -d

run_tests:
	go test -v -cover ./...

migrate-up:
	migrate -path ./db/migration -database 'postgres://api_tester:testing@0.0.0.0:5436/postApi?sslmode=disable' up

migrate-down:
	migrate -path ./db/migration -database 'postgres://api_tester:testing@0.0.0.0:5436/postApi?sslmode=disable' down
