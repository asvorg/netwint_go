package helpers

import (
    "fmt"
    "net"
)

// ResolveDomain resolves a domain name to its IP address.
func ResolveDomain(domain string) (string, error) {
    ips, err := net.LookupIP(domain)
    if err != nil {
        return "", fmt.Errorf("failed to resolve domain %s: %v", domain, err)
    }

    if len(ips) > 0 {
        return ips[0].String(), nil
    }

    return "", fmt.Errorf("no IP addresses found for domain %s", domain)
}
