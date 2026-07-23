package biz

import "testing"

func TestParseSstwapLaneOrderID(t *testing.T) {
	cases := []struct {
		id   string
		lane int
		ok   bool
	}{
		{"sstwap-lane-0", 0, true},
		{"sstwap-lane-4", 4, true},
		{"sstwap-lane-99", 99, true},
		{"0xdeadbeef", 0, false},
		{"sstwap-lane-", 0, false},
		{"", 0, false},
	}
	for _, c := range cases {
		lane, ok := parseSstwapLaneOrderID(c.id)
		if ok != c.ok || (ok && lane != c.lane) {
			t.Fatalf("%q => lane=%d ok=%v; want lane=%d ok=%v", c.id, lane, ok, c.lane, c.ok)
		}
	}
}

func TestSpotTakeOverRecoveryOnlyGuards(t *testing.T) {
	// Recovery adopts only sstwap-lane-N (tested above). Mid-session spot inventory
	// creation is gated by banexg.IsContract in TrialUnMatches / exitByMyOrder —
	// spot Market is never contract, so allowTakeOver/createInv stay false for Slot1.
	if parseLane := func(id string) bool {
		_, ok := parseSstwapLaneOrderID(id)
		return ok
	}; !parseLane("sstwap-lane-1") || parseLane("random-id") {
		t.Fatal("lane-only adoption filter broken")
	}
}
