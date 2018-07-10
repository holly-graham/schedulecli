package schedule

const (
	Monday    = "Monday"
	Tuesday   = "Tuesday"
	Wednesday = "Wednesday"
	Thursday  = "Thursday"
	Friday    = "Friday"
	Saturday  = "Saturday"
	Sunday    = "Sunday"
)

var (
	mondayActivities    []Activity
	tuesdayActivities   []Activity
	wednesdayActivities []Activity
	thursdayActivities  []Activity
	fridayActivities    []Activity
	saturdayActivities  []Activity
	sundayActivities    []Activity
)

type Activity struct {
	Time        string
	Description string
}

func (a Activity) String() string {
	return a.Time + ": " + a.Description
}

func AddActivity(activity Activity, day string) {
	switch day {
	case Monday:
		mondayActivities = append(mondayActivities, activity)
	case Tuesday:
		tuesdayActivities = append(tuesdayActivities, activity)
	case Wednesday:
		wednesdayActivities = append(wednesdayActivities, activity)
	case Thursday:
		thursdayActivities = append(thursdayActivities, activity)
	case Friday:
		fridayActivities = append(fridayActivities, activity)
	case Saturday:
		saturdayActivities = append(saturdayActivities, activity)
	case Sunday:
		sundayActivities = append(sundayActivities, activity)
	}

}

func ListActivities(day string) []Activity {
	switch day {
	case Monday:
		return mondayActivities
	case Tuesday:
		return tuesdayActivities
	case Wednesday:
		return wednesdayActivities
	case Thursday:
		return thursdayActivities
	case Friday:
		return fridayActivities
	case Saturday:
		return saturdayActivities
	case Sunday:
		return sundayActivities
	}

	return nil
}
