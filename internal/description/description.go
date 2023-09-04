package description

import (
	"fmt"
	"time"
)

// Describe ingests a time and returns a friendly
// and human-readable description. If the time is
// in the future, the description will allude so.
func Describe(then time.Time) string {
	now := time.Now()
	duration := now.Sub(then)
	durationAbs := duration.Abs()

	if durationAbs >= month {
		var prefix string
		if now.Before(then) {
			prefix = "on "
		}
		if durationAbs <= 3*month {
			return prefix + then.Format("Jan 2")
		}
		return prefix + then.Format("01/02/2006")
	}

	description := describe(durationAbs)

	if now.Before(then) && durationAbs >= 20*time.Second {
		description = "in " + description
	}
	return description
}

// DescribeVerbosely is similar to Describe, but
// the final description is more verbose. The
// time in question is described more naturally.
func DescribeVerbosely(then time.Time) string {
	now := time.Now()
	duration := now.Sub(then)

	description := describeVerbosely(duration.Abs())

	if now.Before(then) {
		return "in " + description
	}
	return description + " ago"
}

func describe(duration time.Duration) string {
	switch {
	case duration < 20*time.Second:
		return "now"
	case duration < time.Minute:
		return fmt.Sprintf("%.0fs", duration.Seconds())
	case duration < time.Hour:
		return fmt.Sprintf("%.0fm", duration.Minutes())
	case duration < 24*time.Hour:
		return fmt.Sprintf("%.0fh", duration.Hours())
	case duration < 7*24*time.Hour:
		return fmt.Sprintf("%.0fd", duration.Hours()/24)
	default:
		return fmt.Sprintf("%.0fw", duration.Hours()/(24*7))
	}
}

func describeVerbosely(duration time.Duration) string {
	switch {
	case duration < 10*time.Second:
		return "just a few seconds"
	case duration < 30*time.Second:
		return "less than 1 minute"
	case duration < minute+15*time.Second:
		return "about 1 minute"
	case duration < time.Hour:
		return singularOrPlural(duration/minute, "minute")
	case duration < day:
		return singularOrPlural(duration/hour, "hour")
	case duration < week:
		return "almost " + singularOrPlural(duration/day, "day")
	case duration < month:
		return "around " + singularOrPlural(duration/week, "week")
	case duration < year:
		return "roughly " + singularOrPlural(duration/month, "month")
	default:
		return "over " + singularOrPlural(duration/year, "year")
	}
}

func singularOrPlural(numTimeUnits time.Duration, timeUnitName string) (description string) {
	description = fmt.Sprintf("%.0f %s", float64(numTimeUnits), timeUnitName)
	if numTimeUnits > 1 { // plural
		description += "s"
	}
	return
}
