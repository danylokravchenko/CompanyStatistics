package companyService

import (
	"errors"
	"github.com/UndeadBigUnicorn/CompanyStatistics/cache"
	"github.com/UndeadBigUnicorn/CompanyStatistics/models"
	"net/http"
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

		// access field and value of incoming data
		field := fields.Field(i)
		value := values.Field(i)

		// parse value to uint64
		integer, err := strconv.Atoi(value.String())
		if err != nil {
			return http.StatusBadRequest, err
		}

		// so we make totally sure that our data will not be less than 0
		if integer < 0 {
			integer = 0
		}
		v := uint64(integer)

		// some validation of incoming data
		if v < 0 {

			return http.StatusBadRequest, errors.New("data cannot be negative")

		} else if v == 0 && field.Name == "ID" {

			return http.StatusBadRequest, errors.New("ID cannot be 0")

		} else {

			if field.Name == "ID" {

				co, err := c.GetCompany(v)
				if err != nil {
					return http.StatusNotFound, err
				}

				// make copy of original company to prevent it modification if something went wrong
				company = co

			} else {

				// access field via fieldname
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

	company.SetUpdateIsNeeded(true)

	c.PutCompany(company)

	return http.StatusOK, errors.New("")

}
