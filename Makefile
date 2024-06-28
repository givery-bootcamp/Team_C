.PHONY: test
test:
	@docker compose -f compose.test.yml up -d
	@docker compose exec test-backend \
		go test -v -p=1 ./test/e2e/e2e_test.go
	@docker compose -f compose.test.yml down

.PHONY: test-cover
test-cover:
	@cd backend; go test -v -p=1 -cover -tags="unit_test" ./... -coverprofile=cover.out.tmp
	@cd backend; grep -Ev "myapp/docs/docs.go|myapp/main.go|myapp/internal/interface/api/router" cover.out.tmp > cover.out
	@cd backend; rm cover.out.tmp
	@cd backend; go tool cover -html=cover.out -o cover.html
	@open ./backend/cover.html

.PHONY: test-cover-e2e
test-cover-e2e:
	@docker compose -f compose.test.yml up -d --build
	@docker compose exec test-backend \
		go test -v -p=1 ./test/e2e/e2e_test.go
	@docker stop test-backend
	@docker compose -f compose.test.yml down
	@cd backend; go tool covdata textfmt -i coverdir -o ./coverdir/profile.txt
	@cd backend; go tool cover -html=./coverdir/profile.txt -o ./coverdir/profile.html

.PHONY: test-unit
test-unit:
	@cd backend; go test -tags="unit_test" ./...
