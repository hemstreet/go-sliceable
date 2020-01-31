package helpers

import (
    "errors"
    "reflect"
)

const (
    outputPointerNotSliceError = "output pointer is not of type slice"
    inputItemsNotSliceError = "item is not of type slice"
)
func interfaceToReflection(things interface{}) (reflectedThings []reflect.Value, err error) {
    items := reflect.ValueOf(things)

    if items.Kind() != reflect.Slice {
        return nil, errors.New(inputItemsNotSliceError)
    }

    for i := 0; i < items.Len(); i++ {
        reflectedThings = append(reflectedThings, items.Index(i))
    }

    return reflectedThings, nil
}

func MapSlice(things interface{}, fn func(thing interface{}) interface{}, outPtr interface{}) error {
    rr, err := interfaceToReflection(things)

    if err != nil {
        return err
    }

    valuePtr := reflect.ValueOf(outPtr)
    value := valuePtr.Elem()

    if value.Kind() != reflect.Slice {
        return errors.New(outputPointerNotSliceError)
    }

    for i := range rr {
        item := rr[i]
        result := fn(item.Interface())

        value.Set(reflect.Append(value, reflect.ValueOf(result)))
    }

    return nil
}

func Reduce(things interface{}, fn func(thing interface{}) int, outPtr *int) error {
    rr, err := interfaceToReflection(things)

    if err != nil {
        return err
    }

    var total int

    for i := range rr {
        item := rr[i]
        result := fn(item.Interface())
        total += result

    }

    valuePtr := reflect.ValueOf(outPtr)
    value := valuePtr.Elem()

    value.Set(reflect.ValueOf(total))

    return nil
}

func Filter(things interface{}, fn func(thing interface{}) bool, outPtr interface{}) error {
    rr, err := interfaceToReflection(things)

    if err != nil {
        return err
    }

    valuePtr := reflect.ValueOf(outPtr)
    value := valuePtr.Elem()

    if value.Kind() != reflect.Slice {
        return errors.New(outputPointerNotSliceError)
    }

    for i := range rr {
        item := rr[i]

        // Only insert the entity into the slice if true is returned from our function
        if fn(item.Interface()) {
            value.Set(reflect.Append(value, item))
        }
    }

    return nil
}
