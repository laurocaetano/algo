package challenges

// https://dmoj.ca/problem/cco07p2
//
// Example:
// Input: [
//
//	[1, 2, 3, 4, 5, 6],
//	[4, 3, 2, 1, 6, 5]
//
// ]
//
// Output:
// True
// https://gist.github.com/laurocaetano/5515d7f92a46e4fba639274b9857db5b
// full submitted solution, including reading from stdin
func HasTwinSnowflakes(snowflakes [][]int) bool {
	twinMap := make(map[int][][]int)

	for _, snowflake := range snowflakes {
		code := snowflakeCode(snowflake)
		similarCodes := twinMap[code]

		for _, similar := range similarCodes {
			twins := areTwins(snowflake, similar)

			if twins {
				return true
			}
		}

		similarCodes = append(similarCodes, snowflake)
		twinMap[code] = similarCodes
	}

	return false
}

func areTwins(snowflake []int, otherSnowflake []int) bool {
	for start := 0; start < 6; start++ {
		if identicalRight(snowflake, otherSnowflake, start) {
			return true
		}

		if identicalLeft(snowflake, otherSnowflake, start) {
			return true
		}
	}

	return false
}

func identicalRight(snowflake []int, otherSnowflake []int, start int) bool {
	for offset := 0; offset < 6; offset++ {
		if snowflake[offset] != otherSnowflake[(start+offset)%6] {
			return false
		}
	}

	return true
}

func identicalLeft(snowflake []int, otherSnowflake []int, start int) bool {
	for offset := 0; offset < 6; offset++ {
		otherSnowflakeIndex := start - offset

		if otherSnowflakeIndex < 0 {
			otherSnowflakeIndex = otherSnowflakeIndex + 6
		}

		if snowflake[offset] != otherSnowflake[otherSnowflakeIndex] {
			return false
		}
	}
	return true
}

func snowflakeCode(snowflake []int) int {
	sum := 0

	for _, v := range snowflake {
		sum = sum + v
	}

	return sum
}
