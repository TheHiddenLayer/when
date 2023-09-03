package description

import (
	"fmt"
	"math"
	"time"
)

func Describe(t time.Time) string {
	currentTime := time.Now()
	duration := currentTime.Sub(t)

	absDiffSeconds := int(math.Abs(duration.Seconds()))
	if absDiffSeconds >= month {
		return t.Format("01/02/2006")
	}
	description := describeAbsDiff(absDiffSeconds)

	inFuture := duration.Seconds() < 0
	if inFuture {
		description = "in " + description
	}
	return description
}

func DescribeVerbosely(t time.Time) string {
	currentTime := time.Now()
	duration := currentTime.Sub(t)

	absDiffSeconds := int(math.Abs(duration.Seconds()))
	description := describeAbsDiffVerbosely(absDiffSeconds)

	inFuture := duration.Seconds() < 0
	if inFuture {
		return "in " + description
	}
	return description + " ago"
}

func describeAbsDiff(secs int) string {
	switch true {
	case secs < 20:
		return "now"
	case secs < minute:
		return fmt.Sprintf("%d%s", secs, "s")
	case secs < hour:
		return fmt.Sprintf("%d%s", secs/minute, "m")
	case secs < day:
		return fmt.Sprintf("%d%s", secs/hour, "h")
	case secs < week:
		return fmt.Sprintf("%d%s", secs/day, "d")
	default:
		return fmt.Sprintf("%d%s", secs/week, "w")
	}
}

func describeAbsDiffVerbosely(secs int) string {
	switch true {
	case secs < 10:
		return "just a few seconds"
	case secs < minute/2:
		return "less than 1 minute"
	case secs < minute+15:
		return "about 1 minute"
	case secs < hour:
		return grammaticalNumber(secs/minute, "minute")
	case secs < day:
		return grammaticalNumber(secs/hour, "hour")
	case secs < week:
		return grammaticalNumber(secs/day, "day")
	case secs < month:
		return "around " + grammaticalNumber(secs/week, "week")
	case secs < year:
		description := "roughly " + grammaticalNumber(secs/month, "month")
		return description
	default:
		return "over " + grammaticalNumber(secs/year, "year")
	}
}

func grammaticalNumber(numTimeUnits int, timeUnitName string) (description string) {
	description = fmt.Sprintf("%d %s", numTimeUnits, timeUnitName)
	if numTimeUnits > 1 {
		description += "s" // plural time unit
	}
	return
}
