.PHONY: setup build clean clean-all all server

GO-OUT := ./bin/app
GO_FILES := cmd/* handlers/*

TAILWIND_DIR := ./tailwind
NPM := pnpm

#TODO: this does not work ahaahhah

ifeq ($(NPM), pnpm)
LOCKFILE = pnpm-lock.yaml
TLWD_CMD = pnpm --prefix $(TAILWIND_DIR) 

else ifeq ($(NPM), npm)
LOCKFILE = package-lock.json
TLWD_CMD = npm --prefix $(TAILWIND_DIR)

else ifeq ($(NPM), bun)
LOCKFILE = bun.lock
TLWD_CMD = cd $(TAILWIND_DIR) && bun

else ifeq ($(NPM), yarn)
LOCKFILE = yarn.lock
TLWD_CMD = yarn --cwd $(TAILWIND_DIR)
endif

all: setup build

# build
views/%.go: views/*.templ
	templ generate
static/style.css: $(tailwind_dir) views/*.go
	$(TLWD_CMD) run gen-css

$(GO-OUT): $(GO-IN)
	go build -o $(GO-OUT) ./cmd
build: static/style.css views/*.go $(GO-OUT)

# setup
$(TAILWIND_DIR)/node_modules: $(TAILWIND_DIR)/$(LOCKFILE)
$(TAILWIND_DIR)/$(LOCKFILE):
	$(TLWD_CMD) install
go.sum:
	go get github.com/a-h/templ
	go get github.com/labstack/echo/v4
	go get github.com/labstack/echo/v4/middleware
	go install github.com/a-h/templ/cmd/templ@latest
setup: go.sum $(TAILWIND_DIR)/node_modules

# clean
clean:
	@rm -rfv static/style.css views/*_templ.go bin tmp 

clean-all: clean
	@rm -rfv $(TAILWIND_DIR)/package-lock.json $(TAILWIND_DIR)/yarn.lock $(TAILWIND_DIR)/pnpm-lock.yaml $(TAILWIND_DIR)/bun.lock $(TAILWIND_DIR)/node_modules go.sum

server:
	$(GO-OUT)
