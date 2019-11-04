package timeParser

import (
	"github.com/UndeadBigUnicorn/CompanyStatistics/config"
	"time"
)

// Parse time presented in string to normal Go Time
func ParseTime(timeToParse string) time.Time {

	if timeToParse == "" {
		return time.Now()
	}

	parsedTime, err := time.Parse(config.GetSetting("timeLayout").(string), timeToParse)
	if err != nil {
		parsedTime, err = time.Parse(config.GetSetting("dateLayout").(string), timeToParse)
	}

	return parsedTime

}


// Format time to general layout
func FormatTime(today time.Time) string {
	return today.Format(config.GetSetting("timeLayout").(string))
}
