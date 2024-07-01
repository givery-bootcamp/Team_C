.PHONY: test
test:
	@docker compose -f compose.test.yml up -d
	@docker compose exec test-backend \
		go test -v -p=1 ./test/e2e/e2e_test.go
	@docker compose -f compose.test.yml down

.PHONY: test-cover
test-cover:
	@cd backend; go test -v -p=1 -cover -tags='unit_test' ./... -args -test.gocoverdir=$$PWD/coverdir/unit
	@docker compose -f compose.test.yml up -d
	@docker compose exec test-backend \
		go test -v -p=1 ./test/e2e/e2e_test.go
	@docker stop test-backend
	@cd backend; go tool covdata textfmt -i coverdir/unit -o ./coverdir/unit/profile.txt
	@cd backend; go tool covdata textfmt -i coverdir/e2e -o ./coverdir/e2e/profile.txt
	@cd backend; go tool cover -html=./coverdir/e2e/profile.txt -o ./coverdir/e2e/profile.html
	@cd backend; go tool covdata merge -i=coverdir/e2e,coverdir/unit -o ./coverdir/merged
	@cd backend; go tool covdata textfmt -i coverdir/merged -o ./coverdir/merged/profile.txt.tmp
	@cd backend; grep -Ev "myapp/docs/docs.go|myapp/main.go|myapp/internal/.+_mock.go" ./coverdir/merged/profile.txt.tmp > ./coverdir/merged/profile.txt
	@cd backend; rm ./coverdir/merged/profile.txt.tmp
	@cd backend; go tool cover -html=./coverdir/merged/profile.txt -o ./coverdir/merged/profile.html
	@docker compose -f compose.test.yml down

.PHONY: test-cover-unit
test-cover-unit:
	@cd backend; go test -v -p=1 -cover -tags="unit_test" ./... -coverprofile=cover.out.tmp
	@cd backend; grep -Ev "myapp/docs/docs.go|myapp/main.go|myapp/internal/interface/api/router|myapp/internal/.+_mock.go" cover.out.tmp > cover.out
	@cd backend; rm cover.out.tmp
	@cd backend; go tool cover -html=cover.out -o cover.html

.PHONY: test-cover-e2e
test-cover-e2e:
	@docker compose -f compose.test.yml up -d --build
	@docker compose exec test-backend \
		go test -v -p=1 ./test/e2e/e2e_test.go
	@docker stop test-backend
	@docker compose -f compose.test.yml down
	@cd backend; go tool covdata textfmt -i coverdir/e2e -o ./coverdir/e2e/profile.txt.tmp
	@cd backend; grep -Ev "myapp/docs/docs.go|myapp/main.go|myapp/internal/.+_mock.go" ./coverdir/e2e/profile.txt.tmp > ./coverdir/e2e/profile.txt
	@cd backend; rm ./coverdir/e2e/profile.txt.tmp
	@cd backend; go tool cover -html=./coverdir/e2e/profile.txt -o ./coverdir/e2e/profile.html

.PHONY: test-unit
test-unit:
	@cd backend; go test -v -tags="unit_test" ./...
