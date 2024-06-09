build:
	docker-compose build post-api

run:
	docker-compose up post-api -d

run_tests:
	go test -v -cover ./...