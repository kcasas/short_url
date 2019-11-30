# What is .PHONY? https://unix.stackexchange.com/a/321947
.PHONY: run_local ensure_linter lint fast-lint

#####################################
# Build, test & lint
#####################################

GOPATH := $(shell go env GOPATH)
HOSTOS := ${shell go env GOHOSTOS}

run_dev:
	go run cmd/web/server.go

ensure_linter:
	hash golangci-lint || curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(GOPATH)/bin v1.21.0

lint:
	make ensure_linter
	golangci-lint run  ./... && echo "Perfect!"

fast-lint:
	make ensure_linter
	golangci-lint run --fast  ./... && echo "Perfect!"

ensure_migrator:
	hash migrate.$(HOSTOS)-amd64 || curl -L https://github.com/golang-migrate/migrate/releases/download/v4.7.0/migrate.$(HOSTOS)-amd64.tar.gz | tar xvz -C $(GOPATH)/bin

migrate:
	make ensure_migrator
	$(GOPATH)/bin/migrate.$(HOSTOS)-amd64 -database "mysql://$(DB_DSN)" -path migrations/ -verbose up

test:
	go test -cover -race ./...