package models

import (
	"database/sql"
	"time"
)


type Company struct {
	ID uint64 `db:"id"`
	Name string `db:"name"`
	TotalLocations uint64 `db:"totallocations"`
	TotalDoctors uint64 `db:"totaldoctors"`
	TotalUsers uint64 `db:"totalusers"`
	TotalInvitations uint64 `db:"totalinvitations"`
	TotalCreatedReviews uint64 `db:"totalcreatedreviews"`
	TotalOpenedReviews uint64 `db:"totalopenedreviews"`
	UpdatedAt sql.NullString `db:"updatedat"`
	TimeUpdatedAt time.Time
	UpdatedBy sql.NullInt64 `db:"updatedby"`
	UpdateIsNeeded bool
}


// Create new company with initialized companyID
func NewCompany(companyID uint64) *Company {
	return &Company {
		ID: companyID,
		UpdateIsNeeded: false,
	}
}

type CompanyMap map[uint64] *Company