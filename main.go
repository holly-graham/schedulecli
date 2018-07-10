package main

import (
	"fmt"
	"time"

	"github.com/holly-graham/schedulecli/schedule"
	"github.com/manifoldco/promptui"
)

const (
	addActivityCmd  = "Add Activity"
	viewScheduleCmd = "View Schedule"
	backCmd         = "Back"
	view            = "View Day"
	addAnother      = "Add to Same Day"
	mainMenu        = "Main Menu"
	weekOverview    = "Week Overview"
	dayOverview     = "Day Overview"
	differentDay    = "Add to a Different Day"
)

func main() {
	for {
		fmt.Println()

		prompt := promptui.Select{
			Label: "Select One",
			Items: []string{
				addActivityCmd,
				viewScheduleCmd,
			},
		}

		_, result, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		switch result {

		case addActivityCmd:
			err := addActivityPrompt()
			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}

		case viewScheduleCmd:
			err := viewSchedule()
			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}
		}

		time.Sleep(500 * time.Millisecond)
	}

}

func addActivityPrompt() error {

	dayPrompt := promptui.Select{
		Label: "Select Day",
		Items: []string{
			schedule.Monday,
			schedule.Tuesday,
			schedule.Wednesday,
			schedule.Thursday,
			schedule.Friday,
			schedule.Saturday,
			schedule.Sunday,
			backCmd,
		},
	}

	_, result, err := dayPrompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return err
	}

	switch result {

	//addTimeDescription that works for each day of the week

	case schedule.Monday:
		err := addTimeDescription(schedule.Monday)
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return err
		}

	case schedule.Tuesday:
		err := addTimeDescription(schedule.Tuesday)
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return err
		}

	case schedule.Wednesday:
		err := addTimeDescription(schedule.Wednesday)
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return err
		}

	case schedule.Thursday:
		err := addTimeDescription(schedule.Thursday)
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return err
		}

	case schedule.Friday:
		err := addTimeDescription(schedule.Friday)
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return err
		}

	case schedule.Saturday:
		err := addTimeDescription(schedule.Saturday)
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return err
		}

	case schedule.Sunday:
		err := addTimeDescription(schedule.Sunday)
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return err
		}

	case backCmd:
		return nil

	}

	time.Sleep(500 * time.Millisecond)
	return nil
}

func addTimeDescription(day string) error {
	for {
		timePrompt := promptui.Prompt{
			Label: "Time",
		}
		time, err := timePrompt.Run()
		if err != nil {
			return err
		}

		Description := promptui.Prompt{
			Label: "Description",
		}

		description, err := Description.Run()
		if err != nil {
			return err
		}

		Activities := schedule.Activity{
			Time:        time,
			Description: description,
		}

		schedule.AddActivity(Activities, day)

		fmt.Println("Added to", day, time, description)

		prompt := promptui.Select{
			Label: "Select One",
			Items: []string{
				addAnother,
				view,
				differentDay,
				mainMenu,
			},
		}

		_, result, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return err
		}

		switch result {

		case view:
			viewDay := schedule.ListActivities(day)
			for _, day := range viewDay {
				fmt.Println(day)
			}

		case addAnother:
			continue

		case differentDay:
			addActivityPrompt()

		case mainMenu:
			return nil
		}

	}

	return nil
}

func viewSchedule() error {

	for {
		fmt.Println()

		prompt := promptui.Select{
			Label: "Select One",
			Items: []string{
				weekOverview,
				dayOverview,
				backCmd,
			},
		}

		_, result, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return err
		}

		switch result {

		case weekOverview:
			err := weekSchedule()
			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return err
			}

		case dayOverview:
			err := daySchedule()
			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return err
			}
		case backCmd:
			return nil
		}

		time.Sleep(500 * time.Millisecond)

	}
}

func weekSchedule() error {

	week := []string{
		schedule.Monday,
		schedule.Tuesday,
		schedule.Wednesday,
		schedule.Thursday,
		schedule.Friday,
		schedule.Saturday,
		schedule.Sunday,
	}

	for _, day := range week {

		fmt.Println(day)
		viewGivenDay(day)
		time.Sleep(250 * time.Millisecond)
	}

	finalMenu := promptui.Select{
		Label: "Select Iy",
		Items: []string{
			weekOverview,
			dayOverview,
			backCmd,
			mainMenu,
		},
	}

	_, result, err := finalMenu.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return err
	}

	switch result {

	case weekOverview:
		err := weekSchedule()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return err
		}

	case dayOverview:
		err := daySchedule()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return err
		}
	case backCmd:
		return nil

	case mainMenu:
		main()
	}
	return nil
}

func daySchedule() error {

	dayPrompt := promptui.Select{
		Label: "Select Day",
		Items: []string{
			schedule.Monday,
			schedule.Tuesday,
			schedule.Wednesday,
			schedule.Thursday,
			schedule.Friday,
			schedule.Saturday,
			schedule.Sunday,
			backCmd,
			mainMenu,
		},
	}

	_, result, err := dayPrompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return err
	}

	switch result {

	//addTimeDescription that works for each day of the week

	case schedule.Monday:
		err := viewGivenDay(schedule.Monday)
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return err
		}

	case schedule.Tuesday:
		err := viewGivenDay(schedule.Tuesday)
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return err
		}

	case schedule.Wednesday:
		err := viewGivenDay(schedule.Wednesday)
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return err
		}

	case schedule.Thursday:
		err := viewGivenDay(schedule.Thursday)
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return err
		}

	case schedule.Friday:
		err := viewGivenDay(schedule.Friday)
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return err
		}

	case schedule.Saturday:
		err := viewGivenDay(schedule.Saturday)
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return err
		}

	case schedule.Sunday:
		err := viewGivenDay(schedule.Sunday)
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return err
		}

	case backCmd:
		return nil

	case mainMenu:
		main()

	}

	return nil
}

func viewGivenDay(day string) error {

	viewDay := schedule.ListActivities(day)
	for _, day := range viewDay {
		fmt.Println(day)
	}

	finalMenu := promptui.Select{
		Label: "Select One",
		Items: []string{
			weekOverview,
			dayOverview,
			backCmd,
		},
	}

	_, result, err := finalMenu.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return err
	}

	switch result {
	case weekOverview:
		err := weekSchedule()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return err
		}

	case dayOverview:
		err := daySchedule()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return err
		}
	case backCmd:
		return nil
	}

	return nil
}
