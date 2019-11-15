package timeconv

import (
	"testing"
	"time"
)

func parse(v string) *time.Time {
	t, _ := time.Parse(time.RFC3339Nano, v)
	return &t
}
func TestFromUnix(t *testing.T) {
	for _, c := range []struct {
		Input    string
		Expected *time.Time
	}{
		{
			Input:    "0",
			Expected: parse("1970-01-01T09:00:00+09:00"),
		},
		{
			Input:    "1573780170",
			Expected: parse("2019-11-15T10:09:30+09:00"),
		},
		{
			Input:    "1573780170.001",
			Expected: parse("2019-11-15T10:09:30.000000001+09:00"),
		},
		{
			Input:    "invalid data",
			Expected: nil,
		},
	} {
		got := FromUnix(c.Input)
		if got == nil && c.Expected == nil {
			continue
		}
		if !got.Equal(*c.Expected) {
			t.Errorf("input=%s, got=%s, expected=%s", c.Input, got, c.Expected)
		}
	}
}

func TestFromDateString(t *testing.T) {
	tt := time.Now()
	for name, f := range formats {
		expected := tt.Format(f)
		got := FromDateString(expected).Format(f)
		if got != expected {
			t.Errorf("format=%s, got=%s, expected=%s", name, got, expected)
		}

	}
	got := FromDateString("invalid")
	if got != nil {
		t.Errorf("invalid string parsed: %s", got)
	}
}
