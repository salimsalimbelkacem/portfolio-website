GO-OUT=./bin/app

css:
	pnpm --prefix ./tailwind run gen-css

templ:
	templ generate

build: css templ
	go build -o $(GO-OUT) ./cmd
