package service

import "testing"

func TestSplitSlot(t *testing.T) {
	cases := []struct {
		input    string
		wantDay  string
		wantTime string
	}{
		{"Pazartesi 09:00-12:00", "Pazartesi", "09:00-12:00"},
		{"Cuma 14:00-17:00", "Cuma", "14:00-17:00"},
		{"GeçersizSlot", "", ""},
	}

	for _, c := range cases {
		day, tr := splitSlot(c.input)
		if day != c.wantDay || tr != c.wantTime {
			t.Errorf("splitSlot(%q) = (%q,%q); beklenen (%q,%q)",
				c.input, day, tr, c.wantDay, c.wantTime)
		}
	}
}

func TestParseTimeRange(t *testing.T) {
	cases := []struct {
		input   string
		wantStart int
		wantEnd   int
	}{
		{"09:00-12:00", 9, 12},
		{"14:00-17:00", 14, 17},
		{"bozuk", -1, -1},
	}

	for _, c := range cases {
		s, e := parseTimeRange(c.input)
		if s != c.wantStart || e != c.wantEnd {
			t.Errorf("parseTimeRange(%q) = (%d,%d); beklenen (%d,%d)",
				c.input, s, e, c.wantStart, c.wantEnd)
		}
	}
}
