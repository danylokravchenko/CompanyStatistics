package dbworker

import (
	"database/sql"
	"github.com/UndeadBigUnicorn/CompanyStatistics/models"
	"time"
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


// Wrap companies update queries in a transaction
func UpdateBatchCompanies(companies []*models.Company) error {

	tx := db.MustBegin()

	for _, company := range companies {

		company.UpdatedAt = sql.NullString {
			String:  time.Now().Format(timeLayout),
			Valid: true,
		}

		tx.NamedExec(`
			UPDATE companies
			SET 
			totallocations = :totallocations,
			totaldoctors = :totaldoctors,
			totalusers = :totalusers,
			totalinvitations = :totalinvitations,
			totalcreatedreviews = :totalcreatedreviews,
			totalopenedreviews = :totalopenedreviews,
			updatedat = :updatedat,
			updatedby = 1
			WHERE id = :id;
		`, company)

	}

	return tx.Commit()

}
