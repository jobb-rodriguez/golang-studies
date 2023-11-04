# Basic Types
```go
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
```

# Declaring a variable
```go
var var_name var_type = value
```

# Short variable declaration
Go will infer the type.

```go
numCars := 10 // inferred to be an integer

temperature := 0.0 // temperature is inferred to be a floating point value because it has a decimal point

var isFunny = true // isFunny is inferred to be a boolean
```

# Type inference
```go
var i int
j := i // j is also an int

i := 42           // int
f := 3.14         // float64
g := 0.867 + 0.5i // complex128
```

# Same Line Declarations
```go
mileage, company := 80276, "Tesla"

// is the same as

mileage := 80276
company := "Tesla"
```

# Type Sizes
```go
int  int8  int16  int32  int64 // whole numbers

uint uint8 uint16 uint32 uint64 uintptr // positive whole numbers

float32 float64 // decimal numbers

complex64 complex128 // imaginary numbers (rare)
```

## Standard Sizes
- int
- uint
- float64
- complex128

> [!NOTE]
> Specify the size only if needed.

## Type Conversion
```go
temperatureInt := 88
temperatureFloat := float64(temperatureInt)
```

# Preferred Default Types
- bool
- string
- int
- uint32
- byte
- rune
- float64
- complex128

> [!NOTE]
> Stray from the default types only if you're super concerned about performance and memory usage.

# Constants
## Regular
```go
const var_name = var_value
```
## Computed
```go
const firstName = "Lane"
const lastName = "Wagner"
const fullName = firstName + " " + lastName
```

# Formatting strings
[Documentation](https://pkg.go.dev/fmt#hdr-Printing)

- %v for value
- %s for strings
- %d for integer in decimal form
- %f for decimals

```go
fmt.Printf("I am %f years old", 10.523)
// I am 10.523000 years old

// The ".2" rounds the number to 2 decimal places
fmt.Printf("I am %.2f years old", 10.523)
// I am 10.53 years old
```

# Conditionals
```go
if height > 6 {
    fmt.Println("You are super tall!")
} else if height > 4 {
    fmt.Println("You are tall enough!")
} else {
    fmt.Println("You are not tall enough!")
}
```

## Comparison Operators
- == equal to
- != not equal to
- < less than
- > greater than
- <= less than or equal to
- >= greater than or equal to