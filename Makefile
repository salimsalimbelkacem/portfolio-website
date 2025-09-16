.PHONY: setup build clean clean-all all server

GO-OUT := ./bin/app
GO_FILES := cmd/* handlers/*

TAILWIND_DIR := ./tailwind
NPM := pnpm

#TODO: this does not work ahaahhah

# ifeq ($(NPM), pnpm)
LOCKFILE = pnpm-lock.yaml
# else ifeq ($(NPM), npm)
# LOCKFILE = package-lock.json
# else ifeq ($(NPM), bun)
# LOCKFILE = package-lock.json
# else ifeq ($(NPM), yarn)
# LOCKFILE = yarn.lock
# endif

all: setup build

# build
views/%.go: views/*.templ
	templ generate
static/style.css: $(TAILWIND_DIR) views/*.go
	#TODO: hardcode every command for every packet manager :P
	$(NPM) --prefix $(TAILWIND_DIR) run gen-css
$(GO-OUT): $(GO-IN)
	go build -o $(GO-OUT) ./cmd
build: static/style.css views/*.go $(GO-OUT)

# setup
$(TAILWIND_DIR)/node_modules: $(TAILWIND_DIR)/$(LOCKFILE)
$(TAILWIND_DIR)/$(LOCKFILE):
	$(NPM) --prefix $(TAILWIND_DIR) i
go.sum:
	go get github.com/a-h/templ
	go get github.com/labstack/echo/v4
	go get github.com/labstack/echo/v4/middleware
	go install github.com/a-h/templ/cmd/templ@latest
setup: go.sum $(TAILWIND_DIR)/node_modules

# clean
clean:
	@rm -rf static/style.css views/*_templ.go bin tmp 

clean-all: clean
	@rm -rf $(TAILWIND_DIR)/$(LOCKFILE) $(TAILWIND_DIR)/node_modules go.sum

server:
	$(GO-OUT)
