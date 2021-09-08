package godb

func mergeMaps(maps ...Param) Param {
    result := make(Param)
    for _, m := range maps {
        for k, v := range m {
            result[k] = v
        }
    }
    return result
}