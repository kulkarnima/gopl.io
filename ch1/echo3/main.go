package main

import (
    "fmt"
    "os"
    "strings"
)

func Echo3(Args [] string) {

    fmt.Println(strings.Join(os.Args[1:], " "))
}

func main() {
    fmt.Println(strings.Join(os.Args[1:], " "))
    Echo3(os.Args)
}
