# [generics](https://go.dev/doc/tutorial/generics)

Each [type parameter](https://go.dev/ref/spec#Type_parameter_declarations) has a type constraint that acts as a kind of meta-type for the type parameter.

```go
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
    var s V
    for _, v := range m {
        s += v
    }
    return s
}

SumIntsOrFloats[string, int64]

type Number interface {
    int64 | float64
}
```

## comparable

booleans, numbers, strings, pointers, channels, arrays of comparable types, structs whose fields are all comparable types
