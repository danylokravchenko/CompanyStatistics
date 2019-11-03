package timeParser

import "time"

const timeLayout = "2006-01-02" + " " + "15:04:05"


// Parse time presented in string to normal Go Time
func ParseTime(timeToParse string) time.Time {

	if timeToParse == "" {
		return time.Now()
	}

	parsedTime, _ := time.Parse(timeLayout, timeToParse)

	return parsedTime

}


// Format time to general layout
func FormatTime(today time.Time) string {
	return today.Format(timeLayout)
}
