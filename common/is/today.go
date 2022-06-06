package is

import "time"

func Today(date time.Time) bool {
	y1, m1, d1 := date.Date()
	y2, m2, d2 := time.Now().Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}
