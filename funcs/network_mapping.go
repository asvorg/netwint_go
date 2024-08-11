package funcs

import (
    "fmt"
    "net"
)

func DiscoverHosts(ipRange string) {
    for i := 1; i <= 254; i++ {
        ip := fmt.Sprintf("%s.%d", ipRange, i)
        conn, err := net.Dial("ip4:icmp", ip)
        if err == nil {
            fmt.Printf("Host found: %s\n", ip)
            conn.Close()
        }
    }
}