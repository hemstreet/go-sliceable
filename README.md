Go Sliceable
===

Install
====

```go
go get github.com/hemstreet/go-sliceable
```

High Level Theory
===
Implementation of Map, Reduce and Filter. Accept a slice, function and output pointer.

Basic implementation are: 
* `MapSlice(sliceInput, mapFunc, outputSlicePtr)`
* `Reduce(sliceInput, reduceFunc, outputPtr)`
* `Filter(sliceInput, filterFunc, outputSlicePtr)`

Examples
===
MapSlice
``` go

       type Foo struct {
           ID   int
       }
   
       // What you have
       var foos = []Foo{...}
   
       var ids []int
   
       err := sliceable.MapSlice(foos, func(thing interface{}) interface{} {
          return thing.(Foo).ID
       }, &ids)
```

Reduce
``` go

    type Foo struct {
        Years int
    }

    var foos = []Foo{...}

    var totalYears int

    err := sliceable.Reduce(foos, func(thing interface{}) int {
        return thing.(Foo).Years
    }, &totalYears)
```

Filter
``` go
    type Foo struct {
        Org string
    }

    var foos = []Foo{...}

    var filteredFoos []Foo

   err := sliceable.Filter(foos, func(thing interface{}) bool {
        return thing.(Foo).Org  == "Enterprise"
    }, &filteredFoos)
```
