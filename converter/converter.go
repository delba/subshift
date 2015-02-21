package converter

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

var (
	Format    = regexp.MustCompile(`(\d{2,}):(\d{2}):(\d{2}),(\d{3})`)
	TimeUnits = [4]time.Duration{
		time.Hour,
		time.Minute,
		time.Second,
		time.Millisecond,
	}
)

func StringToDuration(s string) (time.Duration, error) {
	var d time.Duration

	matches := Format.FindStringSubmatch(s)[1:]

	for i, match := range matches {
		n, err := strconv.Atoi(match)
		if err != nil {
			return d, err
		}

		d += time.Duration(n) * TimeUnits[i]
	}

	return d, nil
}

func DurationToString(d time.Duration) string {
	var h, m, s, ms time.Duration

	divMod := func(x, y time.Duration) (time.Duration, time.Duration) {
		return x / y, x % y
	}

	h, d = divMod(d, time.Hour)
	m, d = divMod(d, time.Minute)
	s, d = divMod(d, time.Second)
	ms, _ = divMod(d, time.Millisecond)

	return fmt.Sprintf("%02d:%02d:%02d,%03d", h, m, s, ms)
}
