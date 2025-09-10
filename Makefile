.PHONY: setup css build clean templ

GO-OUT := ./bin/app
TAILWIND_DIR := ./tailwind
GO_FILES := cmd/* handlers/*

templ: view/*.templ
view/%.templ:
	templ generate

css:
	pnpm --prefix $(TAILWIND_DIR) run gen-css

build: css templ $(GO-OUT)
$(GO-OUT):
	go build -o $(GO-OUT) ./cmd


setup: go.sum $(TAILWIND_DIR)/node_modules

$(TAILWIND_DIR)/node_modules: $(TAILWIND_DIR)/pnpm-lock.yaml

$(TAILWIND_DIR)/pnpm-lock.yaml:
	pnpm --prefix $(TAILWIND_DIR) i
go.sum:
	go get github.com/a-h/templ
	go get github.com/labstack/echo/v4
	go get github.com/labstack/echo/v4/middleware
	go install github.com/a-h/templ/cmd/templ@latest

clean:
	rm -rf static/style.css views/*_templ.go go.sum bin tmp $(TAILWIND_DIR)/pnpm-lock.yaml $(TAILWIND_DIR)/node_modules
