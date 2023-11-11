GOCMD=go

run: ## Start application
	$(GOCMD) run cmd/main.go

deps: ## Install dependencies
	# go get $(go list -f '{{if not (or .Main .Indirect)}}{{.Path}}{{end}}' -m all)
	$(GOCMD) get -u -t -d -v ./...
	$(GOCMD) mod tidy
	$(GOCMD) mod vendor

swag: ## Generate swagger docs
	swag init -g cmd/main.go   
# 	swag init -g pkg/api/handler/admin.go -o ./cmd/api/docs

nodemon:
	nodemon --exec go run cmd/main.go --signal SIGTERM