package main

import (
    "fmt"
    "sync"
    "argos/internal/ping"
    "argos/internal/config"
)

func main() {
    filePath := "../internal/config/hosts.txt"
	hosts := config.LoadHosts(filePath)

    var wg sync.WaitGroup
    
    for _, host := range hosts {
        wg.Add(1)
        go func(h string) {
            defer wg.Done()
            output := ping.CheckHost(h)
            fmt.Printf("%s → %s\n - latency: %s\n\n", h, output.Status, output.Latency)
        }(host)
    }
    
    wg.Wait()
}