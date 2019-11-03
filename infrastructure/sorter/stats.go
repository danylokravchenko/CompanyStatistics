package sorter

import (
	"github.com/UndeadBigUnicorn/CompanyStatistics/models"
	"sort"
)


// Sort stats by given order
func SortDetailStats(stats []models.UserStats, order string) []models.UserStats {

	switch order {
	case "opened":
		sort.Sort(StatsOpenedSorter(stats))
		break
	case "created":
		sort.Sort(StatsCreatedSorter(stats))
		break
	case "names":
		sort.Sort(StatsNameSorter(stats))
		break
	}
	return stats
}

// StatsSorter sorts by the number of opened reviews
type StatsOpenedSorter []models.UserStats

func (a StatsOpenedSorter) Len() int           { return len(a) }
func (a StatsOpenedSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a StatsOpenedSorter) Less(i, j int) bool { return a[i].Opened > a[j].Opened }


// StatsSorter sorts by the number of created reviews
type StatsCreatedSorter []models.UserStats

func (a StatsCreatedSorter) Len() int           { return len(a) }
func (a StatsCreatedSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a StatsCreatedSorter) Less(i, j int) bool { return a[i].Created > a[j].Created }


// StatsSorter sorts by the name of doctor
type StatsNameSorter []models.UserStats

func (a StatsNameSorter) Len() int           { return len(a) }
func (a StatsNameSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a StatsNameSorter) Less(i, j int) bool { return a[i].Name < a[j].Name }

