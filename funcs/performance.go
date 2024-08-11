package funcs

import (
    "bytes"
    "fmt"
    "os/exec"
    "regexp"
    "strconv"
    "time"
)

// PerformPing performs multiple pings and returns metrics including RTT, jitter, and packet loss.
func PerformPing(ip string, count int) (avgRTT time.Duration, jitter float64, packetLoss float64) {
    var rtts []time.Duration
    var lostPackets int

    for i := 0; i < count; i++ {
        // Execute the ping command
        cmd := exec.Command("ping", "-c", "1", ip)
        var out bytes.Buffer
        cmd.Stdout = &out
        err := cmd.Run()

        if err != nil {
            lostPackets++
            continue
        }

        // Extract RTT from the command output
        output := out.String()
        rtt, err := extractRTT(output)
        if err != nil {
            lostPackets++
            continue
        }

        rtts = append(rtts, rtt)
    }

    if len(rtts) > 0 {
        avgRTT = sumDurations(rtts) / time.Duration(len(rtts))
        jitter = calculateJitter(rtts)
    }

    packetLoss = float64(lostPackets) / float64(count) * 100

    return avgRTT, jitter, packetLoss
}

// extractRTT extracts the RTT value from the ping command output.
func extractRTT(output string) (time.Duration, error) {
    re := regexp.MustCompile(`time=(\d+\.?\d*) ms`)
    match := re.FindStringSubmatch(output)
    if len(match) < 2 {
        return 0, fmt.Errorf("RTT not found in output")
    }
    rttMs, err := strconv.ParseFloat(match[1], 64)
    if err != nil {
        return 0, err
    }
    return time.Duration(rttMs * float64(time.Millisecond)), nil
}

// sumDurations returns the sum of a slice of time.Durations.
func sumDurations(durations []time.Duration) time.Duration {
    var sum time.Duration
    for _, d := range durations {
        sum += d
    }
    return sum
}

// calculateJitter calculates the jitter (variation) from a slice of RTT measurements.
func calculateJitter(rtts []time.Duration) float64 {
    if len(rtts) < 2 {
        return 0
    }

    var totalJitter float64
    var previousRTT time.Duration

    for i, rtt := range rtts {
        if i == 0 {
            previousRTT = rtt
            continue
        }
        jitter := float64(rtt - previousRTT)
        if jitter < 0 {
            jitter = -jitter
        }
        totalJitter += jitter
        previousRTT = rtt
    }

    return totalJitter / float64(len(rtts)-1)
}
