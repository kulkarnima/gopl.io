// Boining prints the boiling point of water
package main

import "fmt"

const boilingF = 212.0

func main() {
    var f = boilingF
    var c = (f - 32) * 5 / 9

    fmt.Printf("boining point = %g deg F or %g deg C\n", f, c)
    // Output:
    // boining point = 212 deg F or 100 deg C
}
