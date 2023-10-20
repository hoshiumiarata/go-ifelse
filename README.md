# ifelse

This package provides equivalent of the ternary operator. It is useful when you want to assign a value to a variable conditionally.

Check out [documentation](https://pkg.go.dev/github.com/hoshiumiarata/go-ifelse) for more information.

## Installation

```bash
go get github.com/hoshiumiarata/go-ifelse
```

## Example

```go
ifelse.If(true, "true", "false") // "true"
ifelse.LazyIf(true, func() int { return "true" }, func() int { return "false" }) // "true"
```
