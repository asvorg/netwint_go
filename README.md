# Network Intelligence

A versatile network intelligence tool written in Go that provides comprehensive functionality for domain resolution, port scanning, service detection, SSL/TLS certificate retrieval, and detailed network performance metrics.

## Features

- **Domain Resolution**: Convert domain names to IP addresses.
- **Ping and Performance Metrics**:
  - **Round Trip Time (RTT)**: Measure the time taken for a packet to travel to the destination and back.
  - **Jitter**: Measure the variation in packet arrival time.
  - **Packet Loss**: Calculate the percentage of lost packets.
- **Port Scanning**:
  - Scan specific ports or a range of ports on an IP address.
  - Supports predefined lists of common ports.
- **Service Detection**: Identify the service running on a specific port of an IP address.
- **SSL/TLS Certificate Retrieval**: Retrieve and display SSL/TLS certificate details for a given IP and port.
- **Network Topology Mapping**: Map open ports and services for a given IP address to visualize network topology.
