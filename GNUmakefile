.DEFAULT_GOAL := help

# Path to the docker-compose binary
DOCKER_COMPOSE_BIN ?= docker compose

.PHONY: help
help: ## Show this help
	@awk 'BEGIN {FS = ":.*?## "} /^[%a-zA-Z0-9_-]+:.*?## / {printf "%-20s%s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: up
up: ## Spin up local testing infrastructure
	$(DOCKER_COMPOSE_BIN) up --wait
	$(DOCKER_COMPOSE_BIN) exec mattermost mmctl --local user create \
		--email test@example.com \
		--username test \
		--password test123 \
		--disable-welcome-email
	$(DOCKER_COMPOSE_BIN) exec mattermost mmctl --local token generate test desc

.PHONY: down
down: ## Destroy local testing infrastructure
	$(DOCKER_COMPOSE_BIN) down

# Manually export MM_LOGIN_ID=test MM_PASSWORD=test123
# Or MM_TOKEN=…
.PHONY: testacc
testacc: ## Run acceptance tests
	TF_ACC=1 MM_URL=http://localhost:8065 go test ./... -v $(TESTARGS) -timeout 2m

.PHONY: test
test: ## Run unit tests
	go test ./... $(TESTARGS)
