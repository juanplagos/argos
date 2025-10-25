package config

import (
	"bufio"
	"strings"
	"log"
	"os"
)

func LoadHosts(filePath string) []string {
	file, err := os.Open(filePath)
    	if err != nil {
        	log.Fatal(err)
    	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var hosts []string

	for scanner.Scan() {
		hostName := strings.TrimSpace(scanner.Text())
		if hostName != "" {
			hosts = append(hosts, hostName)
		}
	}
	return hosts
}