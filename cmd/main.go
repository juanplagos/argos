package main

import (
    "fmt"
    "argos/internal/ping"
    "argos/internal/config"
)

func main() {
    filePath := "../internal/config/hosts.txt"
	hosts := config.LoadHosts(filePath)

    for _, host := range hosts {
        output := ping.CheckHost(host)
        fmt.Printf("%s → %s\n - latency: %s\n\n", host, output.Status, output.Latency)
    }
}