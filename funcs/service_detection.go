package funcs

import (
    "bufio"
    "fmt"
    "net"
    "time"
)

// DetectService tries to connect to a port and identify the service based on the initial response.
func DetectService(ip string, port int) string {
    address := fmt.Sprintf("%s:%d", ip, port)
    
    // Set a timeout for the connection attempt
    conn, err := net.DialTimeout("tcp", address, 5*time.Second)
    if err != nil {
        return "Closed or Unreachable"
    }
    defer conn.Close()

    // Create a buffered reader to read the response
    reader := bufio.NewReader(conn)
    
    // Set a read deadline to avoid hanging on slow responses
    conn.SetReadDeadline(time.Now().Add(5 * time.Second))
    
    // Attempt to read the response
    response, err := reader.Peek(2048) // Read up to 2048 bytes
    if err != nil {
        return "Response Too Large or Error"
    }

    // Convert response to string
    responseStr := string(response)
    
    // Detect common services based on known response patterns
    switch {
    case responseStr[:3] == "220":
        return "FTP (File Transfer Protocol)"
    case responseStr[:6] == "HTTP/1":
        return "HTTP (HyperText Transfer Protocol)"
    case responseStr[:4] == "SSH-":
        return "SSH (Secure Shell)"
    case responseStr[:4] == "220 ":
        return "SMTP (Simple Mail Transfer Protocol)"
    case responseStr[:5] == "RDP ":
        return "RDP (Remote Desktop Protocol)"
    case responseStr[:6] == "ES-1.0":
        return "Elasticsearch"
    default:
        return "Unknown Service"
    }
}
