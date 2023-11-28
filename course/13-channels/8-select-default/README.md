# Tickers

- [time.Tick()](https://golang.org/pkg/time/#Tick) is a standard library function that returns a channel that sends a value on a given interval.
- [time.After()](https://golang.org/pkg/time/#After) sends a value once after the duration has passed.
- [time.Sleep()](https://golang.org/pkg/time/#Sleep) blocks the current goroutine for the specified amount of time.

# Read-only Channels
```go
func readCh(ch <-chan int) {
  // ch can only be read from
  // in this function
}
```

# Write-only Channels
```go
func writeCh(ch chan<- int) {
  // ch can only be written to
  // in this function
}
```