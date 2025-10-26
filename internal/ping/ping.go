package ping

import (
	"os/exec"
	"regexp"
	"strings"
)

type Connection struct {
	Status string
	Latency string
}

func CheckHost(host string) Connection {
	output, err := exec.Command("ping", "-c", "1", host).CombinedOutput()
	parsed_out := string(output)

	var latency string

	latencyPattern := regexp.MustCompile(`time=([0-9]*\.?[0-9]+\s?ms)`)
	if m := latencyPattern.FindStringSubmatch(parsed_out); len(m) > 1 {
		latency = m[1]
	}

	if err != nil || !strings.Contains(parsed_out, "1 received") {
		return Connection{Status: "down", Latency: ""}
	} else {
	return Connection{Status: "up", Latency: latency}
	}
}