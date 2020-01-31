Go Sliceable
===

Implementation of Map, Reduce and Filter. Accept a slice, function and output pointer.

Basic implementation are: 
* `MapSlice(sliceInput, mapFunc, outputSlicePtr)`
* `Reduce(sliceInput, reduceFunc, outputPtr)`
* `Filter(sliceInput, filterFunc, outputSlicePtr)`

MapSlice
``` go

       type Foo struct {
           ID   int
       }
   
       // What you have
       var food = []Foo{...}
   
       var ids []int
   
       if err := MapSlice(foos, func(thing interface{}) interface{} {
          return thing.(Foo).ID
       }, &ids); err != nil {
           t.Errorf("Received error %+v", err)
       }
```

Reduce
``` go

    type Foo struct {
        Years int
    }

    var foos = []Foo{...}

    var totalYears int

    if err := Reduce(foos, func(thing interface{}) int {
        return thing.(Foo).Years
    }, &totalYears); err != nil {
        t.Errorf("Received error %+v", err)
    }
```

Filter
``` go
    type Foo struct {
        Org string
    }

    // What you have
    var foos = []Foo{
       ...
    }

    var filteredFoos []Foo

    if err := Filter(foos, func(thing interface{}) bool {
        return thing.(Foo).Org  == "Enterprise"
    }, &filteredFoos); err != nil {
        t.Errorf("Received error %+v", err)
    }
```