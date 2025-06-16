.PHONY: benchmark
benchmark:
	go test -bench=.

.PHONY: build-docker
build-docker:
	docker build --platform linux/amd64,linux/arm64 . -t quineglobal/blog-quine

.PHONY: publish-docker
publish-docker:
	docker push quineglobal/blog-quine

.PHONY: cover
cover:
	go tool cover -html=cover.out

.PHONY: lint
lint:
	golangci-lint run

.PHONY: start
start:
	go run ./cmd/app

.PHONY: test
test:
	go test -coverprofile=cover.out -shuffle on ./...
