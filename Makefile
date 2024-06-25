.PHONY: test
test:
	@docker compose -f compose.test.yml up -d
	@docker compose exec test-backend \
		go test -v -p=1 ./test/e2e/e2e_test.go
	@docker compose -f compose.test.yml down
