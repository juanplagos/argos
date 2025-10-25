package ping

import (
	"testing"
)

func TestPing(t *testing.T) {
	t.Run("when host is reachable", func(t *testing.T) {

		result := CheckHost("google.com")
		expect := "up"

		if result.Status != expect {
			t.Errorf("got %q, expected %q", result.Status, expect)
		}
	})
	t.Run("when host is unreachable", func(t *testing.T) {

		result := CheckHost("invalid.host")
		expect := "down"

		if result.Status != expect {
			t.Errorf("got %q, expected %q", result.Status, expect)
		}
	})
}