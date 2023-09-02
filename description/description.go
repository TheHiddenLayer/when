package description

import (
    "time"
    "fmt"
    "math"
)

func GenerateDescription(t time.Time) string {
    currentTime := time.Now()
    duration := currentTime.Sub(t)
    absoluteDuration := time.Duration(int(math.Abs(float64(duration.Seconds()))))

    var description string

    if duration.Seconds() < 0 {
        // In the future
        if absoluteDuration < 60 {
            description = "less than a minute from now"
        } else if absoluteDuration < 3600 {
            description = fmt.Sprintf("%d minutes from now", absoluteDuration/60)
        } else if absoluteDuration < 86400 {
            description = fmt.Sprintf("%d hours from now", absoluteDuration/3600)
        } else if absoluteDuration < 604800 {
            description = fmt.Sprintf("%d days from now", absoluteDuration/86400)
        } else if absoluteDuration < 2419200 {
            description = fmt.Sprintf("almost %d weeks from now", absoluteDuration/604800)
        } else if absoluteDuration < 29030400 {
            description = fmt.Sprintf("around %d months from now", absoluteDuration/2419200)
        } else {
            years := absoluteDuration / 29030400
            if years > 1 {
                description = fmt.Sprintf("over %d years from now", years)
            } else {
                description = "almost a year from now"
            }
        } 
    } else {
        // In the past
        if absoluteDuration < 10 {
            description = "just now"
        } else if absoluteDuration < 60 {
            description = "less than a minute ago"
        } else if absoluteDuration < 3600 {
            description = fmt.Sprintf("%d minutes ago", absoluteDuration/60)
        } else if absoluteDuration < 86400 {
            description = fmt.Sprintf("%d hours ago", absoluteDuration/3600)
        } else if absoluteDuration < 604800 {
            description = fmt.Sprintf("%d days ago", absoluteDuration/86400)
        } else if absoluteDuration < 29030400 {
            description = fmt.Sprintf("around %d months ago", absoluteDuration/2419200)
        } else {
            years := absoluteDuration / 29030400
            if years > 1 {
                description = fmt.Sprintf("over %d years ago", years)
            } else {
                description = "almost a year ago"
            }
        } 
    }

    return description
}

