package statsService

import (
	"errors"
	"github.com/UndeadBigUnicorn/CompanyStatistics/cache"
	"github.com/UndeadBigUnicorn/CompanyStatistics/infrastructure/mapReduce"
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
		user.UpdateIsNeeded = true
		stats.TimeMap[today][userID] = user

	} else {
		user := models.UserStats {
			StatsID: 0,
			CompanyID: companyID,
			ID:        userID,
			Name:      userName,
			Today:     timeParser.FormatTime(today),
			TimeToday: today,
			UpdateIsNeeded: true,
		}
		if target == "opened" {
			user.Opened++
		} else {
			user.Created++
		}
		stats.TimeMap[today][userID] = user
	}

	c.PutStatsForCompany(stats)

	return errors.New("")

}


// Get detailed statistic for company filtered by period of time and ordered by
func GetDetailStats(c *cache.Cache, companyID uint64, dateFrom, dateTo time.Time, order string) ([]models.UserStats, error) {

	stats, err := c.GetStatsForCompany(companyID)
	if err != nil {
		return nil, err
	}

	// 1) get detail stats for given period of time
	// 2) intersect timeMaps, compute total stats for users
	// 3) convert maps into arrays and apply sorting by order
	res := mapReduce.MapReduce(mapReduce.FilterStatsMapper(), mapReduce.FilterStatsReducer(order), mapReduce.FilterStatsGenerateInput(stats.TimeMap, dateFrom, dateTo))

	return res.([]models.UserStats), errors.New("")

}