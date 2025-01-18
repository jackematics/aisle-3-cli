package command

import (
	"aisle-3-cli/jobs"
	"fmt"
	"strconv"

	"github.com/urfave/cli/v2"
)

func List(cCtx *cli.Context) error {
	fmt.Println("")
	fmt.Println("Pending jobs: ")
	fmt.Println("")

	for _, room := range jobs.JobsList.Rooms {
		fmt.Println(room.Name + ": ")
		fmt.Println("")
		fmt.Printf("%-4s %-35s %-20s %-20s\n", "", "Job", "Last Completed", "Overdue By")
		fmt.Println("")
		for _, job := range room.Jobs {
			daysOverdueBy := job.DaysOverdueBy()
			if daysOverdueBy < 0 {
				continue
			}

			fmt.Printf("%-4s %-35s %-20s %-20s\n", "", job.Job, job.LastCompleted.ToISOString(), strconv.Itoa(daysOverdueBy) + " days")
		}
		fmt.Println("")
	}
	
	return nil
}