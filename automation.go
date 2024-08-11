package main

import (
    "bufio"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net"
    "os"
    "strconv"
    "strings"
    "netwint_go/funcs"
    "netwint_go/helpers"
)

type Task struct {
    Type   string   `json:"type"`
    Domain string   `json:"domain,omitempty"`
    IP     string   `json:"ip,omitempty"`
    Count  int      `json:"count,omitempty"`
    Ports  []int    `json:"ports,omitempty"`
    Port   int      `json:"port,omitempty"`
}

type Config struct {
    Tasks []Task `json:"tasks"`
}

const defaultConfig = `{
    "tasks": [
        {
            "type": "resolve",
            "domain": "example.com"
        },
        {
            "type": "ping",
            "ip": "192.168.1.1",
            "count": 5
        },
        {
            "type": "scan",
            "ip": "192.168.1.1",
            "ports": [22, 80, 443]
        },
        {
            "type": "service",
            "ip": "192.168.1.1",
            "port": 80
        }
    ]
}`

// Function to check if a domain is valid
func isDomain(domain string) bool {
    return strings.Contains(domain, ".")
}

// Resolve domain to IP address
func resolveDomain(domain string) (string, error) {
    ips, err := net.LookupIP(domain)
    if err != nil {
        return "", err
    }
    return ips[0].String(), nil
}

// Check if the configuration file exists, if not, create it
func loadConfig(filename string) (Config, error) {
    if _, err := os.Stat(filename); os.IsNotExist(err) {
        fmt.Printf("Configuration file %s does not exist. Creating a new one with default settings.\n", filename)
        err := ioutil.WriteFile(filename, []byte(defaultConfig), 0644)
        if err != nil {
            return Config{}, fmt.Errorf("error creating default config file: %w", err)
        }
        fmt.Println("Default configuration file created successfully.")
    }

    file, err := ioutil.ReadFile(filename)
    if err != nil {
        return Config{}, err
    }

    var config Config
    err = json.Unmarshal(file, &config)
    if err != nil {
        return Config{}, err
    }

    return config, nil
}

// Execute a specific task
func executeTask(task Task) {
    if isDomain(task.IP) {
        ip, err := resolveDomain(task.IP)
        if err != nil {
            fmt.Printf("Error resolving domain %s: %s\n", task.IP, err)
            return
        }
        task.IP = ip
    }

    switch task.Type {
    case "resolve":
        ip, err := helpers.ResolveDomain(task.Domain)
        if err != nil {
            fmt.Printf("Error resolving domain %s: %s\n", task.Domain, err)
        } else {
            fmt.Printf("Domain %s resolves to IP: %s\n", task.Domain, ip)
        }

    case "ping":
        avgRTT, jitter, packetLoss := funcs.PerformPingRoutine(task.IP, task.Count)
        fmt.Printf("Ping to %s:\nAverage RTT: %.2f ms\nJitter: %.2f ms\nPacket Loss: %.2f%%\n", task.IP, avgRTT.Seconds()*1000, jitter, packetLoss)

    case "scan":
        funcs.ScanPorts(task.IP, task.Ports)

    case "service":
        service := funcs.DetectService(task.IP, task.Port)
        fmt.Printf("Service on port %d: %s\n", task.Port, service)

    default:
        fmt.Printf("Unknown task type: %s\n", task.Type)
    }
}

// Prompt the user for input and return the provided tasks
func promptForTasks() ([]Task, error) {
    reader := bufio.NewReader(os.Stdin)
    var tasks []Task

    for {
        fmt.Print("Enter task type (resolve, ping, scan, service) or 'done' to finish: ")
        taskType, _ := reader.ReadString('\n')
        taskType = strings.TrimSpace(taskType)

        if taskType == "done" {
            break
        }

        var task Task
        task.Type = taskType

        switch taskType {
        case "resolve":
            fmt.Print("Enter domain to resolve: ")
            domain, _ := reader.ReadString('\n')
            task.Domain = strings.TrimSpace(domain)
            tasks = append(tasks, task)

        case "ping":
            fmt.Print("Enter IP to ping: ")
            ip, _ := reader.ReadString('\n')
            task.IP = strings.TrimSpace(ip)
            fmt.Print("Enter number of pings: ")
            var count int
            fmt.Scan(&count)
            task.Count = count
            tasks = append(tasks, task)

        case "scan":
            fmt.Print("Enter IP to scan: ")
            ip, _ := reader.ReadString('\n')
            task.IP = strings.TrimSpace(ip)
            fmt.Print("Enter ports to scan (comma-separated): ")
            ports, _ := reader.ReadString('\n')
            portList := strings.Split(strings.TrimSpace(ports), ",")
            for _, p := range portList {
                port, _ := strconv.Atoi(p)
                task.Ports = append(task.Ports, port)
            }
            tasks = append(tasks, task)

        case "service":
            fmt.Print("Enter IP to check service: ")
            ip, _ := reader.ReadString('\n')
            task.IP = strings.TrimSpace(ip)
            fmt.Print("Enter port to check: ")
            var port int
            fmt.Scan(&port)
            task.Port = port
            tasks = append(tasks, task)

        default:
            fmt.Println("Invalid task type.")
        }
    }

    return tasks, nil
}

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: automation [config_file]")
        return
    }

    configFile := os.Args[1]
    config, err := loadConfig(configFile)
    if err != nil {
        fmt.Printf("Error loading config: %s\n", err)
        return
    }

    // Prompt user for tasks and add them to the config
    userTasks, err := promptForTasks()
    if err != nil {
        fmt.Printf("Error prompting for tasks: %s\n", err)
        return
    }
    config.Tasks = append(config.Tasks, userTasks...)

    // Execute all tasks from the configuration
    for _, task := range config.Tasks {
        executeTask(task)
    }
}
