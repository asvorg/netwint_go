package testers

import (
    "fmt"
    "netwint_go/helpers"
)

// TestDomainResolution resolves a list of domains and prints their IP addresses.
func TestDomainResolution() {
    domains := []string{
        "www.google.com",
        "www.facebook.com",
        "www.apple.com",
        "www.microsoft.com",
        "www.amazon.com",
        "www.reddit.com",
        "www.wikipedia.org",
        "www.twitter.com",
        "www.linkedin.com",
        "www.netflix.com",
        "www.github.com",
        "www.stackoverflow.com",
    }

    for _, domain := range domains {
        ip, err := helpers.ResolveDomain(domain)
        if err != nil {
            fmt.Println("Error resolving domain:", domain, "-", err)
        } else {
            fmt.Println(domain, "resolves to IP:", ip)
        }
    }
}
