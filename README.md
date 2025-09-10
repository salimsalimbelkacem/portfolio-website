i don't want to justify what is this project and why, it is nothing for the moment,
only me trying out golang and htmx

# dependencies
- `golang`
- `npm` or `pnpm`
- `make`

> make sure to have the latest version of go

now these are optional and you can build the project without them, they are just nice
to have for development
- [air](https://github.com/air-verse/air) 
- a procfile-based process manager like [foreman](https://github.com/ddollar/foreman)
or [overmind](https://github.com/DarthSim/overmind) 

# Setup and build
using `make` will install the rest of the dependencies and libraries in the
project (echo, templ and tailwind) then build everything

using `make setup` will just install the dependencies

# build
you can build everything using `make build`

