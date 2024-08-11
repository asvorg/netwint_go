package helpers

import (
    "net"
    "os/exec"
    "regexp"
    "strconv"
    "strings"
    "fmt"
)

type Helpers struct{}

// ResolveDomain resolves a domain name to an IP address.
func (h *Helpers) ResolveDomain(domain string) (string, error) {
    ips, err := net.LookupIP(domain)
    if err != nil || len(ips) == 0 {
        return "", fmt.Errorf("could not resolve the domain: %v", err)
    }
    return ips[0].String(), nil
}

// Ping sends a ping request to the specified IP and returns the round-trip times.
func (h *Helpers) Ping(ip string) ([]float64, error) {
    out, err := exec.Command("ping", "-c", "4", ip).Output()
    if err != nil {
        return nil, fmt.Errorf("ping command failed: %v", err)
    }

    output := string(out)
    timeRegex := regexp.MustCompile(`time=(\d+\.\d+) ms`)
    matches := timeRegex.FindAllStringSubmatch(output, -1)

    if matches == nil {
        return nil, fmt.Errorf("no ping time found")
    }

    var times []float64
    for _, match := range matches {
        time, err := strconv.ParseFloat(match[1], 64)
        if err == nil {
            times = append(times, time)
        }
    }

    return times, nil
}
