GOCMD=go

run: ## Start application
	$(GOCMD) run cmd/main.go

deps: ## Install dependencies
	# go get $(go list -f '{{if not (or .Main .Indirect)}}{{.Path}}{{end}}' -m all)
	$(GOCMD) get -u -t -d -v ./...
	$(GOCMD) mod tidy
	$(GOCMD) mod vendor

swag: ## Generate swagger docs
	swag init -g internal/infrastructure/api/server.go
# 	swag init -g pkg/api/handler/admin.go -o ./cmd/api/docs # -o is to define the output location of swagger docs folder

nodemon:
	nodemon --exec go run cmd/main.go --signal SIGTERMd

docker-b:
	docker build -t fb-api .

docker: docker-b
	docker ps -q --filter "name=fb-api" | xargs -r docker stop
	docker ps -aq --filter "name=fb-api" | xargs -r docker rm
	docker run -d --name fb-api -p 9000:8989 fb-api