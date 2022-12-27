package challenges

import (
	"testing"
)

func TestHasTwinSnowflakes(t *testing.T) {
	withSnowFlakes := [][]int{
		{1, 2, 3, 4, 5, 6},
		{4, 3, 2, 1, 6, 5},
	}

	withoutSnowflakes := [][]int{
		{1, 2, 3, 4, 5, 6},
		{1, 2, 3, 5, 4, 6},
	}

	found := HasTwinSnowflakes(withSnowFlakes)

	if !found {
		t.Fatal("Expected twin snowflakes to be found, got none")
	}

	found = HasTwinSnowflakes(withoutSnowflakes)

	if found {
		t.Fatal("Expected twin snowflakes to be found, got none")
	}
}
