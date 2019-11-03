package statsService

import (
	"errors"
	"github.com/UndeadBigUnicorn/CompanyStatistics/cache"
	"github.com/UndeadBigUnicorn/CompanyStatistics/infrastructure/timeParser"
	"github.com/UndeadBigUnicorn/CompanyStatistics/models"
	"time"
)


// Update stats for company and/or user (Not Add, because once we add company to cache, we created empty stats for it)
func UpdateStats(c *cache.Cache, companyID, userID uint64, userName, target string, today time.Time) error {

	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	stats, err := c.GetStatsForCompany(companyID)
	if err != nil {
		return err
	}

	if stats.Contains(userID, today) {
		user := stats.TimeMap[today][userID]

		if target == "opened" {
			user.Opened++
		} else {
			user.Created++
		}

	} else {
		users := models.UserStatsMap{}
		user := models.UserStats{
			ID:        userID,
			Name:      userName,
			Today:     timeParser.FormatTime(today),
			TimeToday: today,
		}
		if target == "opened" {
			user.Opened++
		} else {
			user.Created++
		}
		users[userID] = user
	}

	c.PutStatsForCompany(stats)

	return errors.New("")

}


// Get detailed statistic for company filtered by period of time and ordered by
func GetDetailStats(c *cache.Cache, companyID uint64, dateFrom, dateTo time.Time, order string) []models.UserStats {

	stats, _ := c.GetStatsForCompany(companyID)

	users := make([]models.UserStatsMap, 0)

	for today := dateFrom; today.Unix() <= dateTo.Unix(); today.AddDate(0,0,1) {
		users = append(users, stats.TimeMap[today])
	}

	//TODO: add mapReduce

	return nil

}