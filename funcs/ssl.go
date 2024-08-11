// funcs/ssl.go
package funcs

import (
    "crypto/tls"
    "fmt"
)

// GetTLSCert retrieves and displays SSL/TLS certificate details for a given IP and port.
func GetTLSCert(ip string, port int) {
    address := fmt.Sprintf("%s:%d", ip, port)
    // Establish a TLS connection
    conn, err := tls.Dial("tcp", address, nil)
    if err != nil {
        fmt.Println("Error connecting:", err)
        return
    }
    // Retrieve the certificates from the connection
    certs := conn.ConnectionState().PeerCertificates
    if len(certs) == 0 {
        fmt.Println("No certificates found.")
        return
    }
    // Print details for each certificate
    for i, cert := range certs {
        fmt.Printf("Certificate %d:\n", i+1)
        fmt.Printf("  Subject: %s\n", cert.Subject)
        fmt.Printf("  Issuer: %s\n", cert.Issuer)
        fmt.Printf("  Expiration: %s\n", cert.NotAfter)
        fmt.Printf("  Not Before: %s\n", cert.NotBefore)
        fmt.Printf("  DNS Names: %v\n", cert.DNSNames)
        fmt.Printf("  IP Addresses: %v\n", cert.IPAddresses)
    }
}
