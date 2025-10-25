package ping

import (
	"strings"
	"os/exec"
)

type result struct {
	Status string
	Latency string
}

func CheckHost(host string) result {
	output, err := exec.Command("ping", "-c", "1", host).CombinedOutput()
	parsed_out := string(output)

	if err != nil || !strings.Contains(parsed_out, "1 received") {
		//fmt.Println(parsed_out)
		return result{Status: "down"}
	} else {
	//fmt.Println(parsed_out)
	return result{Status: "up"}
	}
}