package cache

import (
	"errors"
	"github.com/UndeadBigUnicorn/CompanyStatistics/dbworker"
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

	for _, company := range dbcompanies {

		companies[company.ID] = &company

	}

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


// watch every 10 seconds for changes in companies map
func (c *Cache) watch() {

	for {

		<-time.After(waitInterval)

		companiesMap := c.GetCompanies()

		companies := make([]*models.Company, 0)

		c.Mutex.Lock()

		for _, company := range companiesMap {

			if company.UpdateIsNeeded() {
				company.SetUpdateIsNeeded(false)
				companies = append(companies, company)
			}

		}

		c.Mutex.Unlock()

		go dbworker.UpdateBatchCompanies(companies)

	}

}