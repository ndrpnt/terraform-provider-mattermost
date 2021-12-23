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
		--email admin@example.com \
		--username admin \
		--password admin \
		--disable-welcome-email
	$(DOCKER_COMPOSE_BIN) exec mattermost mmctl --local token generate admin 'For local testing'

.PHONY: down
down: ## Destroy local testing infrastructure
	$(DOCKER_COMPOSE_BIN) down

.PHONY: testacc
testacc: export TF_ACC=1
testacc: export MM_URL=http://localhost:8065
testacc: export MM_LOGIN_ID=admin
testacc: export MM_PASSWORD=admin
testacc: ## Run acceptance tests
	go test ./... -v $(TESTARGS) -timeout 2m

.PHONY: test
test: ## Run unit tests
	go test ./... $(TESTARGS)
