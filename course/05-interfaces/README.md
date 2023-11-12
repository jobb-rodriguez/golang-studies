# Naming Your Interface Arguments
Naming your arguments increase readaility and clarity.

```go
type Copier interface {
  Copy(sourceFile string, destinationFile string) (bytesCopied int)
}
```

# Print the Variable Type
```go
fmt.Printf("%T\n", v)
```

# Clean Interfaces
## Keep Interfaces Small

## Interface Should Have No Knowledge of Satisfying Types
```go
type car interface {
	Color() string
	Speed() int
	// IsFiretruck() bool
}

type firetruck interface {
	car
	HoseLength() int
}
```

## Interfaces Are Not Classes

[Additional Material](https://blog.boot.dev/golang/golang-interfaces/)