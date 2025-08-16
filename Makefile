css:
	pnpm --prefix ./tailwind run gen-css

templ:
	templ generate

build: css templ
	go build -o ./bin/app ./cmd
