package command

import (
	"aisle-3-cli/isodate"
	"aisle-3-cli/jobs"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

func Complete(cCtx *cli.Context) error {
	completedJobs := cCtx.StringSlice("jobs")
 
	var room *jobs.Room
	for i := range jobs.JobsList.Rooms {
		if jobs.JobsList.Rooms[i].Name == cCtx.String("room") {
			room = &jobs.JobsList.Rooms[i]
		}
	}

	if room == nil {
		panic("invalid room name: " + cCtx.String("room"))
	}

	fmt.Println("Completed Jobs")
	fmt.Println("")

	for i := range completedJobs  {
		for j := range room.Jobs {
			if room.Jobs[j].Job == completedJobs[i] {
				room.Jobs[j].LastCompleted = isodate.ISODate(time.Now())
				fmt.Println("Room: " + room.Name)
				fmt.Println("Job completed: " + room.Jobs[j].Job)
				fmt.Println("Completion date: " + time.Now().Format("2006-01-02"))
				fmt.Println("")
				break
			}
		}
	}

	jsonBytes, err := json.MarshalIndent(jobs.JobsList, "", " ")
	if err != nil {
		panic("Error marshalling json")
	}

	err = os.WriteFile("./jobs.json", jsonBytes, 0644)
	if err != nil {
		panic("Error writing to file")
	}

	return nil
}