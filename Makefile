
fmt:
	go fmt ./...

lint:
	golangci-lint version
	golangci-lint run -v --color always --out-format colored-line-number

ci/lint: export GO111MODULE=on
ci/lint: export GOPROXY=https://goproxy.io,direct
ci/lint: export GOOS=linux
ci/lint: export CGO_ENABLED=0
ci/lint: lint