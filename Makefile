SHELL := /bin/sh
FRONTEND_PORT ?= 5173
BACKEND_PORT ?= 8080
GO_PACKAGES := ./cmd/... ./internal/... ./pkg/...

.PHONY: help install-hooks dev dev-frontend dev-backend build data test test-integration smoke lint fmt pages-preview docker-build docker-push release compose-up compose-down clean hooks-pre-commit hooks-commit-msg hooks-pre-push

help:
	@printf '%s\n' \
		'help              list all targets' \
		'install-hooks     wire .githooks' \
		'dev               run frontend and backend locally' \
		'build             build frontend into docs/ and backend binary' \
		'data              no-op for Mode C' \
		'test              run unit tests' \
		'test-integration  run integration tests' \
		'smoke             run local smoke test' \
		'lint              run linters' \
		'fmt               autoformat' \
		'pages-preview     serve docs/ like GitHub Pages' \
		'docker-build      build linux/amd64 backend image' \
		'docker-push       push backend image to GHCR' \
		'release           tag and publish release artifacts' \
		'compose-up        start local compose stack' \
		'compose-down      stop local compose stack' \
		'clean             remove generated local files'

install-hooks:
	git config core.hooksPath .githooks
	chmod +x .githooks/*

dev:
	$(MAKE) -j2 dev-backend dev-frontend

dev-frontend:
	npm run dev -- --host 127.0.0.1 --port $(FRONTEND_PORT)

dev-backend:
	CGO_ENABLED=0 MPS_API_ADDR=:$(BACKEND_PORT) go run ./cmd/server

build:
	npm run build
	CGO_ENABLED=0 go build -trimpath -o bin/mps-server ./cmd/server
	test -f docs/index.html

data:
	@echo 'Mode C uses runtime jobs; no static data pipeline.'

test:
	npm test -- --run
	CGO_ENABLED=0 go test $(GO_PACKAGES)

test-integration:
	CGO_ENABLED=0 go test -tags=integration ./test/integration/...

smoke:
	./scripts/smoke.sh

lint:
	npm run lint
	npm run typecheck
	CGO_ENABLED=0 go vet $(GO_PACKAGES)

fmt:
	npm run format
	gofmt -w cmd internal pkg test

pages-preview:
	npx http-server docs -p 4173 -c-1

docker-build:
	docker buildx build --platform linux/amd64 -t ghcr.io/$${GITHUB_OWNER:-local}/musician-production-suite:latest .

docker-push:
	docker buildx build --platform linux/amd64 --push -t ghcr.io/$${GITHUB_OWNER}/musician-production-suite:latest .

release:
	./scripts/release.sh

compose-up:
	docker compose -f deploy/docker-compose.yml -f deploy/docker-compose.dev.yml up -d --build

compose-down:
	docker compose -f deploy/docker-compose.yml -f deploy/docker-compose.dev.yml down

clean:
	rm -rf bin coverage dist tmp docs/assets docs/index.html docs/404.html docs/manifest.webmanifest docs/sw.js

hooks-pre-commit:
	.githooks/pre-commit

hooks-commit-msg:
	.githooks/commit-msg .git/COMMIT_EDITMSG

hooks-pre-push:
	.githooks/pre-push
