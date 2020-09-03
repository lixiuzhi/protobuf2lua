package pb2lua

import (
    "os"
    "strings"
    "unicode"
)

func PathExists(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil {
        return true, nil
    }
    if os.IsNotExist(err) {
        return false, nil
    }
    return false, err
}


func getUpperName(name string) string{
    s:=name
    s=s[:1]
    s= strings.Map(unicode.ToUpper, s) + name[1:]
    return s
}