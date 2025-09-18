# My portfolio website
I know you couldn't believe your eyes when you saw my website and wanted to see the
code for yourself, and it is nothing to be ashamed of.

# dependencies
- `golang`
- `pnpm`, `npm`, `yarn` or `bun`
- `make`

> make sure to have the latest version of go

These are optional and you can build the project without them, they are just nice
to have for development
- [air](https://github.com/air-verse/air) 
- a procfile-based process manager like [foreman](https://github.com/ddollar/foreman)
or [overmind](https://github.com/DarthSim/overmind) 

# Setup and build
Using `make` will install the rest of the dependencies and libraries in the
project (echo, templ and tailwind) then build everything

Using `make setup` will only install the dependencies

Using `make build` will only build the project

You can use `make` to only build the project if the setup is already done

# Node Package Manager
By default tailwind will be installed using pnpm so
if you want to use `npm`, `yarn` or `bun` for setup and build:
`make NPM=<your-package-manager>`

> Only the ones specified are suported
