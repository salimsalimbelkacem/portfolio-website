GO-OUT=./bin/app

TAILWIND=./tailwind/
CSS-OUT=static/style.css

templ: view/*_templ.go
view/*_templ.go:
	templ generate


css: $(CSS-OUT)
$(CSS-OUT):
	pnpm --prefix $(TAILWIND) run gen-css


build: css templ $(GO-OUT)
$(GO-OUT):
	go build -o $(GO-OUT) ./cmd


setup: go.sum tailwind/node_modules tailwind/pnpm-lock.yaml

tailwind/node_modules: tailwind/pnpm-lock.yaml
tailwind/pnpm-lock.yaml:
	pnpm --prefix $(TAILWIND) i

go.sum:
	go get github.com/a-h/templ
	go install github.com/a-h/templ/cmd/templ@latest
	go get github.com/labstack/echo/v4
	go get github.com/labstack/echo/v4/middleware

clean:
	rm -rf static/style.css views/*_templ.go go.sum bin tmp tailwind/pnpm-lock.yaml tailwind/node_modules

