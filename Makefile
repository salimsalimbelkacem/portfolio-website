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


setup:
	pnpm --prefix $(TAILWIND) i
	go get github.com/a-h/templ
	go get github.com/labstack/echo/v4
	go get github.com/labstack/echo/v4/middleware@v4.13.4
	go install github.com/a-h/templ/cmd/templ@latest

clean:
	rm -rf static/style.css views/*_templ.go go.sum bin tmp tailwind/pnpm-lock.yaml tailwind/node_modules/

