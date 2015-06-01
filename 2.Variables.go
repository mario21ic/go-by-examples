package main

import "fmt"

func main() {
    var a string = "initial"
    fmt.Println(a)

    // You can declare multiple variables at once
    var b, c int = 1, 2
    fmt.Println(b, c)

    var d = true
    fmt.Println(d)

    // Variables declared without a corresponding initialization are zero-valued
    var e int
    fmt.Println(e)

    // Short form for declaring and initializing a variable
    f := "short"
    fmt.Println(f)
}
