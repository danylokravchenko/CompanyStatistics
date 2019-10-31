package companyService

import (
	"errors"
	"github.com/UndeadBigUnicorn/CompanyStatistics/cache"
	"github.com/UndeadBigUnicorn/CompanyStatistics/models"
	"reflect"
	"strconv"
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


// Update existing company
func UpdateCompany(c *cache.Cache, data interface{}) (int, error) {

	// iterate over request data
	fields := reflect.TypeOf(data)
	values := reflect.ValueOf(data)

	num := fields.NumField()

	var company *models.Company

	for i := 0; i < num; i++ {

		field := fields.Field(i)
		value := values.Field(i)

		v, err := strconv.ParseUint(value.String(), 10, 64)
		if err != nil {
			return 400, err
		}

		if v < 0 {

			return 400, errors.New("data cannot be negative")

		} else if v == 0 && field.Name == "ID" {

			return 400, errors.New("ID cannot be 0")

		} else {

			if field.Name == "ID" {

				var err error
				company, err = c.GetCompany(v)
				if err != nil {
					return 404, err
				}

			} else {

				f := reflect.ValueOf(company).Elem().FieldByName(field.Name)

				// change value using reflection
				if f.IsValid() {

					if f.CanSet() {

						if f.Kind() == reflect.Uint64 {
							if !f.OverflowUint(v) {
								f.SetUint(f.Uint() + v)
							}
						}

					}

				}

			}

		}
	}

	c.PutCompany(company)

	return 200, errors.New("")

}
