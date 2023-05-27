## --------
## Makefile
## --------

help: ## Show this help.
	@sed -ne '/@sed/!s/## //p' $(MAKEFILE_LIST)

run: ## Run the application passing ARGS to the application
	go run main.go $(ARGS)

build: ## Build the application binary
	go build -o lib main.go

lint: ## Run the linter
	golint -set_exit_status ./...

test: ## Run the tests
	go test -v ./...
