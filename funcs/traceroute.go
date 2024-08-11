package funcs

import (
    "bytes"
    "os/exec"
    "regexp"
)

// Traceroute performs a traceroute to a given IP address.
func Traceroute(ip string) ([]string, error) {
    cmd := exec.Command("traceroute", ip)
    var out bytes.Buffer
    cmd.Stdout = &out
    err := cmd.Run()
    if err != nil {
        return nil, err
    }

    return parseTracerouteOutput(out.String()), nil
}

// parseTracerouteOutput parses the output of the traceroute command.
func parseTracerouteOutput(output string) []string {
    var hops []string
    lines := regexp.MustCompile("\n").Split(output, -1)
    for _, line := range lines {
        if line != "" {
            hops = append(hops, line)
        }
    }
    return hops
}
