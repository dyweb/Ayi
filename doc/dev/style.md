# Coding style

Copied from [benchhub](https://github.com/benchhub/benchhub/blob/master/doc/style.md)

TODO: might have a repo for unified coding style guide across project and languages

- https://github.com/golang/go/wiki/CodeReviewComments
- https://golang.org/doc/effective_go.html

## Folder structure

See [directory.md](../directory.md)

## Import

- group import, ref https://github.com/jaegertracing/jaeger/blob/master/CONTRIBUTING.md#imports-grouping
  - std
  - third party
  - packages in project's `lib` folder, they will become third party eventually
  - internal packages
- use `pb` as alias for imported `bhpb` package
- use `i` prefix for package from `go.ice`, i.e. `igrpc "github.com/at15/go.ice/ice/transport/grpc"`
- use `dlog` for `github.com/dyweb/gommon/log` since log is used as package var for package level logger

## Naming

conventions

- `w` for `io.Writer` and similar
- `h` for handler, and sometimes for http client

anti-patterns

- use `GetXXX` when define an interface for data container, so the struct can use `XXX`
- use `pkg.Pkg` as less as possible, but if there is no good name, using `scheduler.Scheduler` for interface is better than `scheduler.Interface` IMO

## Error handling

- return early, if there are many checks at the start of a function, put it in a new function
- use [dyweb/gommon/errors](https://github.com/dyweb/gommon/tree/master/errors)
- DO NOT use https://github.com/pkg/errors, IDE might import it and code compiles because some dependencies are using it
- [ ] TODO
  - [ ] multi error
  - [ ] error wrap