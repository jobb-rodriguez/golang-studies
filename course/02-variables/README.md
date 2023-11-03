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