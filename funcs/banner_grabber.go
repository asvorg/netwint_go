package funcs

import (
    "bufio"
    "fmt"
    "net"
    "time"
)

// GrabBanner connects to a given IP and port and returns the banner (initial response) from the service.
func GrabBanner(ip string, port int) (string, error) {
    address := fmt.Sprintf("%s:%d", ip, port)
    conn, err := net.DialTimeout("tcp", address, 5*time.Second)
    if err != nil {
        return "", fmt.Errorf("error connecting to %s: %v", address, err)
    }
    defer conn.Close()

    // Read the initial response from the service
    conn.SetReadDeadline(time.Now().Add(5 * time.Second))
    reader := bufio.NewReader(conn)
    banner, err := reader.ReadString('\n')
    if err != nil {
        return "", fmt.Errorf("error reading banner: %v", err)
    }

    return banner, nil
}
