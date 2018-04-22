## `gg` - a tool to wrap `go generate`

_Very much work in progress._

```bash
go get -u myitcv.io/gg

# list all go:generate directives in packages ./...
gg -l ./...

# run go generate according to the "algorithm" described below on packages ./...
# here, immutableGen generates code that contains go:generate directives
gg -typed stringer -untyped sortGen,immutableGen,keyGen ./...
```

`gg` was born out of the following scenario:

* it's a good idea to clean all generated files as part of a CI build and regenerate; therefore we need a simple,
  reliable means to re-run `go generate` (or similar) on an entire repo of packages
* some `go generate` programs will generate code that itself contains `go generate` directives; this requires `go generate`
  to be called multiple times before a "fixed point" is reached
* some `go generate` programs do type checking (e.g. [`stringer`](https://godoc.org/golang.org/x/tools/cmd/stringer));
  let's call these **typed generators** (vs **untyped generators**)
* typed generators often (always?) fail in situations where a package does not compile
* we therefore need to ensure our untyped generators run first and repeatedly until there are no more changes (if we
  assume that it is generally generated code from the untyped generators that allows a package to otherwise compile)
* then we can run the typed generators; it there is any change, we need to rinse and repeat with the untyped generators,
  then the typed generators... until we reach a fixed point with the typed generators

Whilst it's possible to achieve all of this on a per-project basis by writing a relatively simple program to wrap things
up, there is some merit in writing a tool to wrap `go generate`:

* the tool can be reused by others
* existing `go generate` programs can be re-used with zero effort (other than needing to classify them as either typed
  or untyped)


## TODO

* The code is very much WIP
* Docs
* Some automatic means of detecting typed vs untyped generators?

## Credit

* https://github.com/rsc/gt
* `go generate` source code in [the main Go repo](https://github.com/golang/go/tree/master/src/cmd/go)

