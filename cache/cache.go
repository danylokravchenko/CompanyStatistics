package cache

import (
	"errors"
	"github.com/UndeadBigUnicorn/CompanyStatistics/dbworker"
	"github.com/UndeadBigUnicorn/CompanyStatistics/infrastructure/mapReduce"
	"github.com/UndeadBigUnicorn/CompanyStatistics/models"
	"github.com/patrickmn/go-cache"
	"sync"
	"time"
)

type Cache struct {
	*cache.Cache
	Mutex *sync.RWMutex
}

const(
	defaultExpiration = 30 * time.Minute
	cleanupInterval   = 60 * time.Minute
	waitInterval      = 10 * time.Minute
	companyMap        = "companies"
	statsMap 		  = "stats"
)


// Init new cache
func New() *Cache {

	c:= &Cache{cache.New(defaultExpiration, cleanupInterval), &sync.RWMutex{}}

	c.loadCompanies()
	go c.watch()

	return c

}


// Load companies from database
func (c *Cache) loadCompanies() {

	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	dbcompanies := dbworker.LoadCompanies()

	companies := make(models.CompanyMap)
	stats := make(models.StatsMap)

	for _, company := range dbcompanies {

		companies[company.ID] = &company
		// 1) collect company stats from database
		// 2) split big array of stats into smaller ones
		// 3) transform into time map
		userStatsArray := dbworker.LoadStats(company.ID)
		companyStats := &models.Stats{
			CompanyID: company.ID,
			TimeMap: models.TimeMap{},
		}

		if len(userStatsArray) > 0 {
			timeMap := mapReduce.MapReduce(mapReduce.LoadStatsMapper(), mapReduce.LoadStatsReducer(), mapReduce.LoadStatsGenerateInput(userStatsArray))
			companyStats.TimeMap = timeMap.(models.TimeMap)
		}

		stats[company.ID] = companyStats

	}

	c.Put(statsMap, &stats)
	c.Put(companyMap, &companies)

}


// Put an item with no expiration
func (c *Cache) Put(key string, value interface{}) {
	c.Set(key, value, cache.NoExpiration)
}


// Get companies map
func (c *Cache) GetCompanies() models.CompanyMap {

	if companiesMap, ok := c.Get(companyMap); ok {
		return *companiesMap.(*models.CompanyMap)
	} else {
		return models.CompanyMap{}
	}

}


// Put companies map
func (c *Cache) PutCompanies(companies *models.CompanyMap) {

	c.Put(companyMap, companies)

}


// Get specific company
func (c *Cache) GetCompany(companyID uint64) (*models.Company, error) {

	companies := c.GetCompanies()

	company, ok := companies[companyID]

	if !ok {
		return nil, errors.New("Company was not found")
	}

	return company, nil

}


// Put company in map
func (c *Cache) PutCompany(company *models.Company) {

	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	companies := c.GetCompanies()

	companies[company.ID] = company

	c.PutCompanies(&companies)

}


// Get companies stats map
func (c *Cache) GetStats() models.StatsMap {

	if stats, ok := c.Get(statsMap); ok {
		return *stats.(*models.StatsMap)
	} else {
		return models.StatsMap{}
	}

}


// Put companies stats map
func (c *Cache) PutStats(stats *models.StatsMap) {

	c.Put(statsMap, stats)

}


// Get specific company stats
func (c *Cache) GetStatsForCompany(companyID uint64) (*models.Stats, error) {

	stats := c.GetStats()

	statistic, ok := stats[companyID]

	if !ok {
		return nil, errors.New("Stats was not found")
	}

	return statistic, nil

}


// Put stats for specific company in map
func (c *Cache) PutStatsForCompany(stats *models.Stats) {

	statistics := c.GetStats()

	statistics[stats.CompanyID] = stats

	c.PutStats(&statistics)

}


// watch every 10 minutes for changes in companies map
func (c *Cache) watch() {

	for {

		<-time.After(waitInterval)

		companiesMap := c.GetCompanies()
		statsMap := c.GetStats()

		c.Mutex.Lock()

		wg := &sync.WaitGroup{}

		wg.Add(2)

		// update companies
		go func(wg *sync.WaitGroup, companiesMap models.CompanyMap) {
			companies := make([]*models.Company, 0)
			for _, company := range companiesMap {
				if company.UpdateIsNeeded {
					company.UpdateIsNeeded = false
					companies = append(companies, company)
				}
			}
			go dbworker.UpdateBatchCompanies(companies)
			wg.Done()
		}(wg, companiesMap)

		// update or create personal stats
		go func(wg *sync.WaitGroup, statsMap models.StatsMap) {
			// 1) Filter personal stats by company
			// 2) Check if update is needed for personal stats
			// 3) Divide into 2 arrays: [0] - to insert to db; [1] - to update in db
			statsInterface := mapReduce.MapReduce(mapReduce.UpdateStatsMapper(), mapReduce.UpdateStatsReducer(), mapReduce.UpdateStatsGenerateInput(statsMap))
			stats := statsInterface.([][]models.UserStats)
			go dbworker.InsertBatchStats(stats[0])
			go dbworker.UpdateBatchStats(stats[1])
			wg.Done()
		}(wg, statsMap)


		wg.Wait()
		c.Mutex.Unlock()

	}

}