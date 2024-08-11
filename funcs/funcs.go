package funcs

import (
    "fmt"
    "os/exec"
    "regexp"
    "strconv"
    "strings"
	)

// PingHost pings the specified host and returns the average response time in milliseconds.
func PingHost(host string) (float64, error) {
    // Execute the ping command
    cmd := exec.Command("ping", "-c", "4", host)
    output, err := cmd.Output()
    if err != nil {
        return 0, fmt.Errorf("failed to ping host %s: %v", host, err)
    }

    // Convert the output to a string and look for the average time line
    outputStr := string(output)
    avgTime, err := extractAvgTime(outputStr)
    if err != nil {
        return 0, fmt.Errorf("failed to extract average time for host %s: %v", host, err)
    }

    return avgTime, nil
}

// extractAvgTime extracts the average ping time from the ping command output.
func extractAvgTime(output string) (float64, error) {
    // Use regex to find the average time in the ping output
    re := regexp.MustCompile(`min/avg/max/mdev = [\d.]+/([\d.]+)/[\d.]+/[\d.]+ ms`)
    matches := re.FindStringSubmatch(output)
    if len(matches) < 2 {
        return 0, fmt.Errorf("could not find average time in ping output")
    }

    // Convert the average time to a float64
    avgTimeStr := strings.TrimSpace(matches[1])
    avgTime, err := strconv.ParseFloat(avgTimeStr, 64)
    if err != nil {
        return 0, fmt.Errorf("failed to parse average time: %v", err)
    }

    return avgTime, nil
}
