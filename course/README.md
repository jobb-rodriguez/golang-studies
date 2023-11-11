# Learn Go from boot.dev
[Reference](https://github.com/bootdotdev/fcc-learn-golang-assets)

# Personal Note
During part 1 and 2, the author focused on coding and taking notes.

From part 3 onwards, the author will focus on coding.

# On Naked Returns
Only use them on short and simple functions

# On Early Returns
Useful in cleaning up code (see example below).

```go
func getInsuranceAmount(status insuranceStatus) int {
  if !status.hasInsurance(){
    return 1
  }
  if status.isTotaled(){
    return 10000
  }
  if !status.isDented(){
    return 0
  }
  if status.isBigDent(){
    return 270
  }
  return 160
}
```