GO-OUT=./bin/app

TAILWIND_DIR=./tailwind/
CSS-OUT=static/style.css

templ: view/*_templ.go
view/*_templ.go:
	templ generate

css: $(CSS-OUT)
$(CSS-OUT):
	pnpm --prefix $(TAILWIND_DIR) run gen-css


build: css templ $(GO-OUT)
$(GO-OUT):
	go build -o $(GO-OUT) ./cmd

setup: go.sum $(TAILWIND_DIR)/node_modules $(TAILWIND)/pnpm-lock.yaml
$(TAILWIND_DIR)/node_modules: $(TAILWIND)/pnpm-lock.yaml
$(TAILWIND_DIR)/pnpm-lock.yaml:
	pnpm --prefix $(TAILWIND_DIR) i
go.sum:
	go get github.com/a-h/templ
	go install github.com/a-h/templ/cmd/templ@latest
	go get github.com/labstack/echo/v4
	go get github.com/labstack/echo/v4/middleware

clean:
	rm -rf static/style.css views/*_templ.go go.sum bin tmp $(TAILWIND_DIR)/pnpm-lock.yaml $(TAILWIND)/node_modules
