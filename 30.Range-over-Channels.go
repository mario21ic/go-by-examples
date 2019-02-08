package main

import "fmt"

func main() {

    // We'll iterate over 3 values in the `queue` channel.
    queue := make(chan string, 3)
    queue <- "first"
    queue <- "second"
    queue <- "third"
    close(queue)

    // This `range` iterates over each element as it's
    // received from `queue`. Because we `close`d the
    // channel above, the iteration terminates after
    // receiving the 2 elements.
    x := 0
    for elem := range queue {
        fmt.Printf("ite: %d\t", x)
        fmt.Println(elem)
        x++
    }
}

