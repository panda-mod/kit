package deepcopy

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	m := map[string]any{
		"a": 1,
		"b": 2,
		"c": map[string]any{
			"d": 3,
			"e": 4,
		},
	}
	m2 := Map(m)
	m["c"].(map[string]any)["d"] = 5
	fmt.Printf("Input:  %v\n", m)
	fmt.Printf("Output: %v\n", m2)
	fmt.Println("---")
}
