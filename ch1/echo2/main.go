package main

import (
    "fmt"
    "os"
)

func Echo2(Args [] string) {
    s, sep := "", ""

    for _, arg := range os.Args[1:] {
        s += sep + arg
        sep = " "
    }
    fmt.Println(s)
}

func main() {

    Echo2(os.Args)
}
