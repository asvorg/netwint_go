package funcs

import (
    "fmt"
    "net"
    "time"
)

// ScanPorts checks if ports are open on a given IP address.
func ScanPorts(ip string, ports []int) {
    // Iterate over the list of ports
    for _, port := range ports {
        address := fmt.Sprintf("%s:%d", ip, port)
        
        // Set a timeout for the connection attempt
        conn, err := net.DialTimeout("tcp", address, 1*time.Second)
        if err != nil {
            // If there's an error, the port is likely closed
            fmt.Printf("Port %d closed\n", port)
            continue
        }
        conn.Close()
        // If no error, the port is open
        fmt.Printf("Port %d open\n", port)
    }
}
	func ScanPortsRange(ip string, minPort int, maxPort int) {
		if minPort < 1 || maxPort > 65535 || minPort > maxPort {
			fmt.Println("Invalid port range. Ports should be between 1 and 65535, and minPort should be less than or equal to maxPort.")
			return
		}
		fmt.Printf("Scanning ports %d to %d on %s...\n", minPort, maxPort, ip)
	
		// Iterate over the range of ports
		for port := minPort; port <= maxPort; port++ {
			address := fmt.Sprintf("%s:%d", ip, port)
			
			// Set a timeout for the connection attempt
			conn, err := net.DialTimeout("tcp", address, 1*time.Second)
			if err != nil {
				// If there's an error, the port is likely closed
				fmt.Printf("Port %d closed\n", port)
				continue
			}
			conn.Close()
			// If no error, the port is open
			fmt.Printf("Port %d open\n", port)
		}
	}
