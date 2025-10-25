package main

import (
    "fmt"
    "argos/internal/ping"
    "argos/internal/config"
)

func main() {
    filePath := "/home/jlagos/argos/internal/config/hosts.txt"
	hosts := config.LoadHosts(filePath)

    for _, host := range hosts {
        output := ping.CheckHost(host)
        fmt.Printf("%s → %s\n", host, output.Status)
    }
}