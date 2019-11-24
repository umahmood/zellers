package zellers

// Day of week
const (
	Saturday int = iota
	Sunday
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
)

// zc zellers formula
func zc(month, day, year int) int {
	if month < 3 {
		month += 12
		year--
	}
	return (day +
		int(((month+1)*13)/5) +
		year +
		int(year/4) -
		int(year/100) +
		int(year/400)) % 7
}

// Congruence returns the day of week the month, day and year corresponds to. It
// is the callers responsibility to make sure month, day and year are valid.
func Congruence(month, day, year int) string {
	d := ""
	c := zc(month, day, year)
	switch c {
	case Saturday:
		d = "Saturday"
	case Sunday:
		d = "Sunday"
	case Monday:
		d = "Monday"
	case Tuesday:
		d = "Tuesday"
	case Wednesday:
		d = "Wednesday"
	case Thursday:
		d = "Thursday"
	case Friday:
		d = "Friday"
	}
	return d
}
