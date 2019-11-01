package dbworker

import "time"

const timeLayout = "2006-01-02" + " " + "15:04:05"


// Parse dbtime presenting in string to normal Go Time
func parseTime(dbtime string) time.Time {

	parsedTime, _ := time.Parse(timeLayout, dbtime)

	return parsedTime

}
