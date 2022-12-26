package challenges

import (
	"reflect"
	"testing"
)

// [3, 2, 5]
// [3, 3, 5]
// [4, 3, 5]
// [4, 4, 5]
// [5, 4, 5]
func TestShortestLines(t *testing.T) {
	expected := []int{5, 4, 5}

	result := ShortestLines(4, []int{3, 2, 5})

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Expected %v, got %v", expected, result)
	}
}
