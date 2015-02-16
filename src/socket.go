package main

 import "fmt"

 func main() {
    for i := 0; i < 10000; i++ {
        defer fmt.Print(i)
    }
    return
 }