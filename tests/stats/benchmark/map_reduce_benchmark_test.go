package benchmark

import (
	"fmt"
	"github.com/UndeadBigUnicorn/CompanyStatistics/infrastructure/mapReduce"
	"github.com/UndeadBigUnicorn/CompanyStatistics/infrastructure/timeParser"
	"github.com/UndeadBigUnicorn/CompanyStatistics/models"
	"testing"
	"github.com/Pallinder/go-randomdata"
	"time"
)

func BenchmarkMapReduce(b *testing.B) {
	b.ReportAllocs()
	b.SetParallelism(4)
	dateFrom := timeParser.ParseTime("2019-08-01")
	dateTo := timeParser.ParseTime("2019-11-03")
	timeMap := models.TimeMap{}

	for j := 0; j < 60; j++ {
		//day := randomdata.FullDateInRange("2019-08-01", "2019-11-03")
		month := randomdata.Number(01, 12)
		day := randomdata.Number(01, 31)
		date := fmt.Sprintf("2019-%d-%d",month, day)
		today := timeParser.ParseTime(date)
		timeMap[today] = models.UserStatsMap{}

		for k := 0; k < 100; k++ {
			id := uint64(randomdata.Number(1, 100))
			timeMap[today][id] = models.UserStats{
				ID:        id,
				Name:      randomdata.FullName(1),
				Created:   uint64(randomdata.Number(0, 1000)),
				Opened:    uint64(randomdata.Number(0, 1000)),
				Today:     "",
				TimeToday: time.Time{},
			}
		}

	}

	for i := 0; i < b.N; i++ {
		mapReduce.MapReduce(mapReduce.FilterStatsMapper(), mapReduce.FilterStatsReducer("opened"), mapReduce.FilterStatsGenerateInput(timeMap, dateFrom, dateTo))
	}

}

