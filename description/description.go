package description

import (
	"fmt"
	"math"
	"time"
)

func Describe(t time.Time) string {
	currentTime := time.Now()
	duration := currentTime.Sub(t)
	absoluteDuration := time.Duration(int(math.Abs(float64(duration.Seconds()))))

	description := absoluteTimeDifference(absoluteDuration)

	futureTime := duration.Seconds() < 0
	if futureTime {
		description += " from now"
	} else {
		description += " ago"
	}

	return description
}

// TODO: correct plurals
func absoluteTimeDifference(absoluteDuration time.Duration) string {
	var description string
	if absoluteDuration < minute {
		description = "less than a minute"
	} else if absoluteDuration < hour {
		description = fmt.Sprintf("%d minutes", absoluteDuration/minute)
	} else if absoluteDuration < day {
		description = fmt.Sprintf("%d hours", absoluteDuration/hour)
	} else if absoluteDuration < week {
		description = fmt.Sprintf("%d days", absoluteDuration/day)
	} else if absoluteDuration < month {
		description = fmt.Sprintf("almost %d weeks", absoluteDuration/week)
	} else if absoluteDuration < year {
		months := absoluteDuration / month
		if months < 1 {
			description = "almost a month"
		} else if month == 1 {
			description = "around a month"
		} else {
			description = fmt.Sprintf("over %d months", months)
		}
	} else {
		years := absoluteDuration / year
		if years < 1 {
			description = "almost a year"
		} else if year == 1 {
			description = "around a year"
		} else {
			description = fmt.Sprintf("over %d years", years)
		}
	}
	return description
}
