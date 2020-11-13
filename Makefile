PROJECT_NAME := "task-repo"
APP_NAME := "service"
PKG := "./"
CMD := "$(PKG)/cmd/$(APP_NAME)"

build: ## Build the binary file
	@go build -o $(APP_NAME).o -i -v $(CMD)