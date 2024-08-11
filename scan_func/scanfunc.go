package scanfunc

import (
    "fmt"
    "net"
    "time"
    "bufio"
    "strings"
)

type ScanFunc struct{}

// CommonPorts defines a list of common ports to scan.
var CommonPorts = [40]uint16{
    80, 443, 21, 22, 25, 110, 143, 53, 3389, 137,
    138, 139, 445, 3306, 5432, 8080, 23, 179, 465,
    587, 636, 993, 995, 1723, 2049, 3268, 3269, 5433,
    5985, 5986, 8081, 8443, 9000, 9090, 9091, 9100,
    9200, 9300, 9418, 27017,
}

// Connect attempts to connect to the specified IP and port.
func (s *ScanFunc) Connect(ip string, port uint16) bool {
    address := fmt.Sprintf("%s:%d", ip, port)
    conn, err := net.DialTimeout("tcp", address, 5*time.Second)
    if err != nil {
        fmt.Printf("Failed to connect to %s\n", address)
        return false
    }
    defer conn.Close()

    response := s.getServerResponse(conn)
    fmt.Printf("Response from %s: %s\n", address, response)
    return true
}