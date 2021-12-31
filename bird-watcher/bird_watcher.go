package birdwatcher

const daysPerWeek = 7

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) (result int) {
	for i := range birdsPerDay {
		result += birdsPerDay[i]
	}
	return
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) (result int) {
	from := (week-1) * daysPerWeek
	to := from + daysPerWeek
	min := min(to, len(birdsPerDay))
	for i := from; i < min; i++ {
		result += birdsPerDay[i]
	}
	return
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := range birdsPerDay {
		if i % 2 == 0 {
			birdsPerDay[i]++
		}
	}
	return birdsPerDay
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
