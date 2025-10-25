package ping

import (
	"testing"
)

func TestPing(t *testing.T) {
	t.Run("when host is reachable", func(t *testing.T) {

		Connection := CheckHost("google.com")
		expect := "up"

		if Connection.Status != expect {
			t.Errorf("got %q, expected %q", Connection.Status, expect)
		}
	})
	t.Run("when host is unreachable", func(t *testing.T) {

		Connection := CheckHost("invalid.host")
		expect := "down"

		if Connection.Status != expect {
			t.Errorf("got %q, expected %q", Connection.Status, expect)
		}
	})
}