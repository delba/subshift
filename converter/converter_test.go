package converter

import "testing"

func TestStringToDuration(t *testing.T) {
	s, _ := StringToDuration("12:34:56,789")

	if s.String() != "12h34m56.789s" {
		t.Error("Expected 12h34m56.789s. Got", s)
	}
}

func TestDurationToString(t *testing.T) {
	d, _ := StringToDuration("12:34:56,789")
	s := DurationToString(d)

	if s != "12:34:56,789" {
		t.Error("Expected 12:34:56,789. Got", s)
	}
}
