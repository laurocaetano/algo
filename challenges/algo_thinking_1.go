package challenges

// https://dmoj.ca/problem/lkp18c2p1
//
// Example:
// Imput =>
//
//	newPeople = 4
//	peoplerPerLine = [3, 2, 5]
//
// Output =>
//
//	[2, 3, 3, 4]
func ShortestLines(newPeople int, peoplePerLine []int) []int {
	for i := 1; i <= newPeople; i++ {
		shortestLine := shortestLineIndex(peoplePerLine)
		peoplePerLine[shortestLine]++
	}

	return peoplePerLine
}

func shortestLineIndex(lines []int) int {
	shortest := 0

	for i, numberOfPeople := range lines {
		if numberOfPeople < lines[shortest] {
			shortest = i
		}
	}

	return shortest
}
