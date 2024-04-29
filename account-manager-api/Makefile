APP_NAME = minha_aplicacao
GO_SRC = ./account-manager-api
BUILD_DIR = ./build
ENV_FILE := .env
CONTAINER_NAME := postgres

define docker_cleanup
	docker stop $(CONTAINER_NAME) || true; \
	docker rm $(CONTAINER_NAME) || true; \
	docker compose down -v
endef

$(ENV_FILE):
	@if [ ! -f $(ENV_FILE) ]; then \
		cp .env.template $(ENV_FILE); \
	fi

.PHONY: build
build:
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(APP_NAME) $(GO_SRC)


.PHONY: prod
run:
	@docker-compose up -d
	@echo "Executando a aplicação Go..."
	./$(BUILD_DIR)/$(APP_NAME)


.PHONY: manage-db
manage-db:
	@if [ -z "$$(docker ps -q -f name=$(CONTAINER_NAME))" ]; then \
		$(call docker_cleanup); \
		docker compose up --build -d; \
	fi


.PHONY: dev
dev: $(ENV_FILE) manage-db
	@npx kill-port 8080
	@go mod tidy
	@air


.PHONY: clean
clean:
	rm -rf $(BUILD_DIR)
	rm -rf go.sum
	rm -rf go.mod
