# Anonymous Structs
Use anonymous structs only when you won't need to use a struct again.

For example, sometimes I'll use one to create the shape of some JSON data in HTTP handlers.

```go
myCar := struct {
  Make string
  Model string
} {
  Make: "tesla",
  Model: "model 3"
}

type car struct {
  Make string
  Model string
  Height int
  Width int
  // Wheel is a field containing an anonymous struct
  Wheel struct {
    Radius int
    Material string
  }
}
```