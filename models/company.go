package models

import "time"


type Company struct {
	ID uint64
	Name string
	IsActive bool
	Settings string
	Templates string
	TotalLocations uint64
	TotalDoctors uint64
	TotalUsers uint64
	TotalInvitations uint64
	TotalCreatedReviews uint64
	TotalOpenedReviews uint64
	CreatedAt time.Time
	CreatedBy uint64
	UpdatedAt time.Time
	UpdatedBy uint64
	DeletedAt time.Time
	DeletedBy uint64
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
		updateIsNeeded: true,
	}
}

type CompanyMap map[uint64] *Company