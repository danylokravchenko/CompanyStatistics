package dbworker

import (
	"github.com/UndeadBigUnicorn/CompanyStatistics/models"
)


// Load companies from database
func LoadCompanies() []models.Company {

	var companies []models.Company
	db.Select(&companies, `
		SELECT id, name, totallocations, totaldoctors, totalusers,
		totalinvitations, totalcreatedreviews, totalopenedreviews, updatedat, updatedby
		FROM companies 
		WHERE deletedat is NULL
	`)

	for _, company := range companies {
		if company.UpdatedAt.Valid {
			company.TimeUpdatedAt = parseTime(company.UpdatedAt.String)
		}
	}

	return companies

}
