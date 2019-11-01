package models

import (
	"database/sql"
	"time"
)


type Company struct {
	ID uint64 `db:"id"`
	Name string `db:"name"`
	//IsActive bool
	//Settings string
	//Templates string
	TotalLocations uint64 `db:"totallocations"`
	TotalDoctors uint64 `db:"totaldoctors"`
	TotalUsers uint64 `db:"totalusers"`
	TotalInvitations uint64 `db:"totalinvitations"`
	TotalCreatedReviews uint64 `db:"totalcreatedreviews"`
	TotalOpenedReviews uint64 `db:"totalopenedreviews"`
	//CreatedAt time.Time
	//CreatedBy uint64
	UpdatedAt sql.NullString `db:"updatedat"`
	TimeUpdatedAt time.Time
	UpdatedBy sql.NullInt64 `db:"updatedby"`
	//DeletedAt time.Time
	//DeletedBy uint64
	updateIsNeeded bool
}


// Getter for updateIsNeeded field, because we don't to store it in database
func (c *Company) UpdateIsNeeded() bool {
	return c.updateIsNeeded
}

// Setter for updateIsNeeded field, because we don't to store it in database
func (c *Company) SetUpdateIsNeeded(status bool) {
	c.updateIsNeeded = status
}


// Create new company with initialized companyID
func NewCompany(companyID uint64) *Company {
	return &Company {
		ID: companyID,
		updateIsNeeded: false,
	}
}

type CompanyMap map[uint64] *Company