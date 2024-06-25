.PHONY: test
test:
	@docker compose -f compose.test.yml up -d
	@docker compose exec test-backend \
		go test -v -p=1 ./test/e2e/e2e_test.go
	@docker compose -f compose.test.yml down

.PHONY: test-cover
test-cover:
	@docker compose -f compose.test.yml up -d
	@docker compose exec test-backend \
		go test -v -p=1 -cover ./... -coverprofile=cover.out
	@docker compose exec test-backend \
		go tool cover -html=cover.out -o cover.html
	@open ./backend/cover.html
	@docker compose -f compose.test.yml down
