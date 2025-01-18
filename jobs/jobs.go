package jobs

import (
	"aisle-3-cli/isodate"
	"encoding/json"
	"os"
	"slices"
	"time"
)



type Job struct {
	Job string `json:"job"`
	Frequency isodate.Frequency `json:"frequency"`
	LastCompleted isodate.ISODate `json:"last-completed"`
}

func (job Job) DaysOverdueBy() int {
	frequencyAdjusted := time.Time(job.LastCompleted).
		Add(time.Duration(24 * 7 * job.Frequency.Weeks) * time.Hour).
		Add(time.Duration(24 * 7 * 4 * job.Frequency.Months) * time.Hour).
		Add(time.Duration(24 * 7 * 4 * 12 * job.Frequency.Years) * time.Hour)

	overdueDuration := time.Now().Sub(frequencyAdjusted)

	return int(overdueDuration.Hours() / 24)
}

type Room struct {
	Name string `json:"name"`
	Jobs []Job `json:"jobs"`
}

type Jobs struct {
	Rooms []Room `json:"rooms"`
}

func getJobs() Jobs {
	file, err := os.Open("jobs.json")
	if err != nil {
		panic("Error opening jobs json")
	}
	defer file.Close()

	var jobs Jobs
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&jobs)
	if err != nil {
		panic("Error decoding jobs json")
	}

	for i := range jobs.Rooms {
		slices.SortFunc(jobs.Rooms[i].Jobs, func(a, b Job) int {
			if a.DaysOverdueBy() > b.DaysOverdueBy() {
				return -1
			} else if a.DaysOverdueBy() < b.DaysOverdueBy() {
				return 1
			}

			return 0
		})
	}

	return jobs
}

var JobsList = getJobs()