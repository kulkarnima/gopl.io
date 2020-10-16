package main

import (
    "fmt"
    "os"
)

func Echo1(Args [] string) {
    var s, sep string

    for i := 1; i < len(os.Args); i++ {
        s += sep + os.Args[i]
        sep = " "
    }
    fmt.Println(s)
}

func main() {

    Echo1(os.Args)
}
