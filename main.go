package main

import (
    "fmt"
    "netwint_go/funcs"
)

func main() {
    host := "www.google.com"

    avgTime, err := funcs.PingHost(host)
    if err != nil {
        fmt.Println("Error pinging host:", err)
    } else {
        fmt.Printf("Average ping time to %s: %.2f ms\n", host, avgTime)
    }
}
