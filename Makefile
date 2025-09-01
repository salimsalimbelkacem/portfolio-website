GO-OUT=./bin/app
TAILWIND=./tailwind/

css:
	pnpm --prefix $(TAILWIND) run gen-css

templ:
	templ generate

build: css templ
	go build -o $(GO-OUT) ./cmd

setup:
	pnpm --prefix $(TAILWIND) i
	go get github.com/a-h/templ
	go get github.com/labstack/echo/v4
	go get github.com/labstack/echo/v4/middleware@v4.13.4
	go install github.com/a-h/templ/cmd/templ@latest
