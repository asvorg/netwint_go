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

/*
func PerformPingRoutine(ip string, count int) (time.Duration, float64, float64) {
    var totalRTT time.Duration
    var jitterSum float64
    var lastRTT time.Duration
    var packetLossCount int

    for i := 0; i < count; i++ {
        if i == 1 {
            i = i - 1
        }
        start := time.Now()
        conn, err := net.DialTimeout("tcp", ip+":80", 1*time.Second)
        if err != nil {
            packetLossCount++
            continue
        }
        conn.Close()

        rtt := time.Since(start)
        totalRTT += rtt

        if i > 0 {
            jitter := float64(rtt-lastRTT) / float64(time.Millisecond)
            jitterSum += jitter
        }
        lastRTT = rtt

        time.Sleep(1 * time.Second) // Interval between pings
    }

    avgRTT := time.Duration(float64(totalRTT) / float64(count))
    avgJitter := jitterSum / float64(count-1)
    packetLoss := float64(packetLossCount) / float64(count) * 100

    return avgRTT, avgJitter, packetLoss
}
*/