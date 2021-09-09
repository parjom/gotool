package goutil

import (
    "reflect"
)

func MergeMaps(maps ...map[string]interface{}) map[string]interface{} {
    result := make(map[string]interface{})
    for _, m := range maps {
        for k, v := range m {
            result[k] = v
        }
    }
    return result
}

func Clone(oldObj interface{}) interface{} {
    newObj := reflect.New(reflect.TypeOf(oldObj).Elem())
    oldVal := reflect.ValueOf(oldObj).Elem()
    newVal := newObj.Elem()
    for i := 0; i < oldVal.NumField(); i++ {
        newValField := newVal.Field(i)
        if newValField.CanSet() {
            newValField.Set(oldVal.Field(i))
        }
    }

    return newObj.Interface()
}