up:
	docker-compose up -d --build

test: ### Run go test
	go test ./...

coverage: ### Run go test whit Coverage
	go test -cover ./...

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help
