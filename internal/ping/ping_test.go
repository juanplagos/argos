package ping

import (
	"testing"
	"regexp"
)

func TestPing(t *testing.T) {
	t.Run("when host is reachable", func(t *testing.T) {

		connection := CheckHost("google.com")
		expect := "up"

		if connection.Status != expect {
			t.Errorf("got %q, expected %q", connection.Status, expect)
		}
	})
	t.Run("when host is unreachable", func(t *testing.T) {

		connection := CheckHost("invalid.host")
		expect := "down"

		if connection.Status != expect {
			t.Errorf("got %q, expected %q", connection.Status, expect)
		}
	})
	t.Run("get latency", func(t *testing.T) {

		connection := CheckHost("google.com")
			latencyPattern := `^\d+(\.\d+)?\s?ms$`
			matched, err := regexp.MatchString(latencyPattern, connection.Latency)
			if err != nil || !matched {
				t.Errorf("expected latency matching pattern %q, got %q",
					latencyPattern,
					connection.Latency)
			}
		})
}