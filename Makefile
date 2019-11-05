PROJECTNAME=$(shell basename "$(PWD)")

GOBASE=$(shell pwd)
GOBIN=$(GOBASE)
GOFILES=$(wildcard *.go)

run:
	@echo "  >  Starting server..."
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go run $(GOBIN) server.go

docker-up:
	@echo "  >  Building docker..."
	docker-compose -f $(GOBIN)/docker-compose.yml up -d --build

docker-down:
	@echo "  >  Making docker down..."
	docker-compose -f $(GOBIN)/docker-compose.yml down

test-stats:
	@echo "  >  Testing stats..."
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) cd $(GOBIN)/tests/stats; go test

test-company:
	@echo "  >  Testing stats..."
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) cd $(GOBIN)/tests/company; go test