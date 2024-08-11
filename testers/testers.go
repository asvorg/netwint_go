package testers

import (
    "fmt"
    "netwint_go/helpers"
    "netwint_go/scan_func"
)

type Testers struct{}

func ResolveAndPing() {
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
        "www.iltalehti.fi",
    }

    helper := helpers.Helpers{}
    scanner := scanfunc.ScanFunc{}

    for _, domain := range domains {
        ip, err := helper.ResolveDomain(domain)
        if err != nil {
            fmt.Printf("Could not resolve the domain name: %s\n", domain)
            continue
        }

        fmt.Printf("%s resolves to IP: %s\n", domain, ip)

        times, err := helper.Ping(ip)
        if err != nil {
            fmt.Printf("Error pinging %s: %s\n", ip, err)
            continue
        }

        fmt.Printf("Ping times to %s: %v\n", ip, times)
        avgTime := avg(times)
        fmt.Printf("Average time: %.2f ms\n\n", avgTime)

        fmt.Printf("Scanning ports on %s...\n", ip)
        scanner.SimpleScanMostCommonPorts(ip)
    }
}

func avg(times []float64) float64 {
    sum := 0.0
    for _, time := range times {
        sum += time
    }
    return sum / float64(len(times))
}
