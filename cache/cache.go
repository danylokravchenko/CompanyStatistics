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
	waitInterval      = 5 * time.Minute
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

	//for {
	//
	//	<-time.After(waitInterval)
	//
	//	users := c.GetUsers()
	//
	//	// run it in gorutine because storing data in db may take some time
	//	// and map with updated users could be out-dated
	//	go func(users models.UserMap) {
	//		for _, user := range users {
	//
	//			go func(user *models.User) {
	//				if user.UpdateIsNeeded {
	//
	//					if dbworker.UserExists(user.UserID) {
	//						// update user in db
	//						go dbworker.UpdateUser(user)
	//
	//						for _, transaction := range user.Transactions {
	//							if !dbworker.TransactionExists(transaction.TransactionID) {
	//								go dbworker.CreateTransaction(&transaction)
	//							}
	//						}
	//
	//						for _, deposit := range user.Deposits {
	//							if !dbworker.DepositExists(deposit.DepositID) {
	//								go dbworker.CreateDeposit(&deposit)
	//							}
	//						}
	//
	//					} else {
	//						// add user to db
	//						go dbworker.CreateUser(user)
	//
	//						for _, transaction := range user.Transactions {
	//							go dbworker.CreateTransaction(&transaction)
	//						}
	//
	//						for _, deposit := range user.Deposits {
	//							go dbworker.CreateDeposit(&deposit)
	//						}
	//
	//					}
	//
	//					user.UpdateIsNeeded = false
	//
	//				}
	//
	//			}(user)
	//
	//		}
	//
	//	}(users)
	//
	//}

}