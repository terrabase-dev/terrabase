ROOT_DIR := $(PWD)

API_DIR           := api
UV_LOCK           := $(API_DIR)/uv.lock
API_LOCK_STAMP    := $(API_DIR)/.locked
API_INSTALL_STAMP := $(API_DIR)/.installed
API_PYPROJECT     := $(API_DIR)/pyproject.toml

DOCKER_CERTS     := docker/certs
DOCKER_BOOTSTRAP := docker/scripts/bootstrap.sh

UV_ENV_PREFIX := if [ -n "$$VIRTUAL_ENV" ]; then export UV_PROJECT_ENVIRONMENT=$$VIRTUAL_ENV; fi &&

.PHONY: dev
dev: $(DOCKER_CERTS)
	docker compose -f docker/docker-compose.yml up --build -d

.PHONY: stop
stop:
	docker compose -f docker/docker-compose.yml down

.PHONY: restart
restart: stop dev

.PHONY: migrate
DATABASE_URL ?= postgres://terrabase:terrabase@localhost:5432/terrabase?sslmode=disable
migrate:
	DATABASE_URL=$(DATABASE_URL) go run ./cmd/bedrockmigrate

.PHONY: proto
proto:
	cd proto && buf generate

.PHONY: install-api
install-api: $(API_INSTALL_STAMP)

$(API_LOCK_STAMP): $(API_PYPROJECT)
	cd $(API_DIR) && \
	uv lock && \
	cd $(ROOT_DIR) && \
	touch $(API_LOCK_STAMP)

$(UV_LOCK): $(API_LOCK_STAMP)

$(API_INSTALL_STAMP): $(API_PYPROJECT) $(UV_LOCK)
	@$(UV_ENV_PREFIX) cd $(API_DIR) && \
	uv sync --all-groups && \
	cd $(ROOT_DIR) && \
	touch $(API_INSTALL_STAMP)

$(DOCKER_CERTS): $(DOCKER_BOOTSTRAP)
	sudo ./$(DOCKER_BOOTSTRAP)
