install-dependencies:
	dep ensure

run: install-dependencies
	export PORT=8001 REDIS_AUTH_PASS=pass
	go build -o redis-go && ./redis-go

test: install-dependencies
	go test
