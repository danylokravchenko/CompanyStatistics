package companyService

import (
	"errors"
	"github.com/UndeadBigUnicorn/CompanyStatistics/cache"
	"github.com/UndeadBigUnicorn/CompanyStatistics/models"
)


// Get company to cache
func AddCompany(c *cache.Cache, companyID uint64) error {

	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	companies := c.GetCompanies()

	_, ok := companies[companyID]

	if ok {
		return errors.New("Company with this id already exists! Try again")
	}

	companies[companyID] = models.NewCompany(companyID)

	c.PutCompanies(&companies)

	return nil

}
