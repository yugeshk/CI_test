package mypack

import "testing"

func TestCheckip(t *testing.T) {
	var table = []struct {
		ip  string
		ans bool
	}{
		{"192.168.1.1", true},
		{"11.1455.23.34", false},
		// {"01.2.2.2", false},
	}
	for _, tuple := range table {
		ip := tuple.ip
		want := tuple.ans

		if got := checkip(ip); got != want {
			t.Errorf("Failed for %s, Got: %v, Want: %v", ip, got, want)
		}
	}
}

func Testclean(t *testing.T) {
	var v string
	v = clean("A@@#?yugesh")
	if v != "Ayugesh" {
		t.Error("Incorrect cleaning")
	}
}
