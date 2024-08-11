package main

import (
    "fmt"
    "os"
    "netwint_go/funcs"
	"netwint_go/helpers"
	"strconv"
)

func printHelp() {
    fmt.Println("Usage: netwint_go [command] [options]")
    fmt.Println("Commands:")
    fmt.Println("  resolve [domain]       Resolve a domain to IP address")
    fmt.Println("  ping [ip] [count]      Ping an IP address and display performance metrics")
    fmt.Println("  scan [ip] [minPort] [maxPort]  Scan ports on an IP address")
    fmt.Println("  service [ip] [port]    Detect service running on a port")
	fmt.Println("  map [ipRange]          Map network by discovering active hosts in the given IP range")
	fmt.Println("  banner [ip] [port]     Grab the banner from a port on the given IP address")
    fmt.Println("  help                   Display this help message")
	
}

func main() {


    if len(os.Args) < 2 {
        printHelp()
        return
    }

    command := os.Args[1]

    switch command {
    case "resolve":
        if len(os.Args) < 3 {
            fmt.Println("Usage: resolve [domain]")
            return
        }
        domain := os.Args[2]
        ip, err := helpers.ResolveDomain(domain)
        if err != nil {
            fmt.Println("Error resolving domain:", err)
            return
        }
        fmt.Println("IP address:", ip)
        
    case "ping":
        if len(os.Args) < 4 {
            fmt.Println("Usage: ping [ip] [count]")
            return
        }
        ip := os.Args[2]
        count, err := strconv.Atoi(os.Args[3])
        if err != nil {
            fmt.Println("Invalid count:", err)
            return
        }
        avgRTT, jitter, packetLoss := funcs.PerformPing(ip, count)
        fmt.Printf("Average RTT: %.2f ms\nJitter: %.2f ms\nPacket Loss: %.2f%%\n", avgRTT.Seconds()*1000, jitter, packetLoss)
        
    case "scan":
        if len(os.Args) < 5 {
            fmt.Println("Usage: scan [ip] [minPort] [maxPort]")
            return
        }
        ip := os.Args[2]
		ports := []int{
			20,    // FTP Data Transfer
			21,    // FTP Command
			22,    // SSH
			23,    // Telnet
			25,    // SMTP
			53,    // DNS
			67,    // DHCP Server
			68,    // DHCP Client
			80,    // HTTP
			110,   // POP3
			143,   // IMAP
			161,   // SNMP
			162,   // SNMP Trap
			194,   // IRC
			443,   // HTTPS
			465,   // SMTPS
			514,   // Syslog
			587,   // SMTP (Submission)
			631,   // IPP (Internet Printing Protocol)
			993,   // IMAPS
			995,   // POP3S
			1080,  // SOCKS Proxy
			1433,  // Microsoft SQL Server
			3306,  // MySQL
			3389,  // RDP (Remote Desktop Protocol)
			5432,  // PostgreSQL
			5900,  // VNC
			6379,  // Redis
			6667,  // IRC
			8080,  // HTTP Alternate
			8443,  // HTTPS Alternate
			9200,  // Elasticsearch
			9300,  // Elasticsearch Transport
			27017, // MongoDB
			5000,  // UPnP
			9000,  // Various services
			49152, // Dynamic/Private Ports
			65535, // Dynamic/Private Ports
		}	
		funcs.ScanPorts(ip, ports)

    case "service":
        if len(os.Args) < 4 {
            fmt.Println("Usage: service [ip] [port]")
            return
        }
        ip := os.Args[2]
        port, _ := strconv.Atoi(os.Args[3])
        service := funcs.DetectService(ip, port)
        fmt.Printf("Service on port %d: %s\n", port, service)

	case "map":
        if len(os.Args) < 3 {
            fmt.Println("Usage: map [ipRange]")
            return
        }
        ipRange := os.Args[2]
        funcs.DiscoverHosts(ipRange)
	
	case "banner":
        if len(os.Args) < 4 {
            fmt.Println("Usage: banner [ip] [port]")
            return
        }
        ip := os.Args[2]
        port, _ := strconv.Atoi(os.Args[3])
        banner, err := funcs.GrabBanner(ip, port)
        if err != nil {
            fmt.Println("Error grabbing banner:", err)
            return
        }
        fmt.Printf("Banner for port %d: %s\n", port, banner)

	case "help":
        printHelp()


    default:
        fmt.Println("Unknown command:", command)
        printHelp()
    }
}
