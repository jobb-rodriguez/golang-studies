# Naming Your Interface Arguments
Naming your arguments increase readaility and clarity.

```go
type Copier interface {
  Copy(sourceFile string, destinationFile string) (bytesCopied int)
}
```