package models

import (
	"time"
)


type Stats struct {
	CompanyID uint64 `db:"companyid"`
	TimeMap TimeMap
}


type UserStats struct {
	StatsID uint64 `db:"statsid" json:"-,omitempty"` // don't need in json
	CompanyID uint64 `db:"companyid" json:"-,omitempty"` // don't need in json
	ID uint64 `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	Created uint64 `db:"created" json:"created"`
	Opened uint64 `db:"opened" json:"opened"`
	Today string `db:"today" json:"today"`
	TimeToday time.Time `json:"-,omitempty"` // don't need in json
	UpdateIsNeeded bool `json:"-,omitempty"` // don't need in json
}


// today date: usersMap
// userID: statistic for this user during this day
type TimeMap map[time.Time] UserStatsMap

type UserStatsMap map[uint64] UserStats


// Check if company stats contains this user
func (s *Stats) Contains(userID uint64, today time.Time) bool {

	if _, ok := s.TimeMap[today]; ok {
		_, contains := s.TimeMap[today][userID]
		return contains
	} else {
		s.TimeMap[today] = UserStatsMap{}
		return false
	}

}

type StatsMap map[uint64] *Stats