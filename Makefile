build: install-dependencies
	cd cmd && go build -o redisgo

install-dependencies:
	dep ensure

run: install-dependencies
	export PORT=8001 REDIS_AUTH_PASS=pass
	cd cmd && go build -o redis-go && ./redis-go

test: install-dependencies
	export PORT=3001 && go test ./...

test-with-report: install-dependencies
	export PORT=3001 && go test ./... -coverprofile=coverage.txt -covermode=atomic

test-html-report: test-with-report
	go tool cover -html=coverage.txt
