package main

import (
    "fmt"
    "netwint_go/funcs"
	"netwint_go/helpers"
)

func main() {
    ip,_ := helpers.ResolveDomain("www.iltalehti.fi")
	fmt.Printf(ip)
	fmt.Printf("\n")
    ports := []int{80, 443, 22, 21, 25, 3306, 5432} // List of ports to scan
    for _, port := range ports {
        service := funcs.DetectService(ip, port)
        fmt.Printf(service)
    }
}
