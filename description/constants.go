package description

const (
	second = 1
	minute = 60 * second
	hour   = 60 * minute
	day    = 24 * hour
	week   = 7 * day
	month  = (4 * week) + (2 * day) // 30 days
	year   = 12 * month
)
