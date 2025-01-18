package jobs

import (
	"aisle-3-cli/isodate"
	"encoding/json"
	"os"
)



type Job struct {
	Job string `json:"job"`
	Frequency isodate.Frequency `json:"frequency"`
	LastCompleted isodate.ISODate `json:"last-completed"`
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

	return jobs
}

var JobsList = getJobs()