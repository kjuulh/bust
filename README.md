# Bust

Bust is a platform agnostic pipeline. It is built on top of `dagger` and
`kjuulh/byg`. The goal of this project is to produce a way to easily extend and
interact with a golang pipeline for CI

## Examples

see `examples` for example pipelines, thought do note that the project usually
needs to be self-contained like the `ci` folder. `Bust` is built with `Bust`
after all.

To run simply `go run example/golang-bin/main.go`, this may require certain
setup for docker-hub, or alternate registries.
