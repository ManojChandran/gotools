package main

import (
    "fmt"
)

func main() {
  fmt.Println(Fib(6))
}

func Fib(n int) int {
        if n < 2 {
                return n
        }
        return Fib(n-1) + Fib(n-2)
}
