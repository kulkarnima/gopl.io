package main

import (
    "fmt"
    "os"
)

func main() {
    for index, arg := range os.Args[1:] {
        fmt.Println(arg, index + 1)
    }
}
