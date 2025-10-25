package ping

import (
	"os/exec"
	"strings"
)

type Connection struct {
	Status string
	Latency string
}

func CheckHost(host string) Connection {
	output, err := exec.Command("ping", "-c", "1", host).CombinedOutput()
	parsed_out := string(output)

	if err != nil || !strings.Contains(parsed_out, "1 received") {
	return Connection{Status: "down"}
	} else {
	return Connection{Status: "up"}
	}
}