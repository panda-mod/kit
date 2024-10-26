package convert

import (
	"fmt"
	"testing"
)

var testCases = []any{
	"true",
	"false",
	"1",
	"0",
	true,
	false,
	1,
	"",
	[]int{1},
}

func TestBool(t *testing.T) {
	for _, tc := range testCases {
		fmt.Printf("Input:  %v\n", tc)
		fmt.Printf("Output: %t\n", Bool(tc))
		fmt.Println("---")
	}
}

func TestString(t *testing.T) {
	for _, tc := range testCases {
		fmt.Printf("Input:  %v\n", tc)
		fmt.Printf("Output: %s\n", String(tc))
		fmt.Println("---")
	}
}

func TestFloat64(t *testing.T) {
	for _, tc := range testCases {
		fmt.Printf("Input:  %v\n", tc)
		fmt.Printf("Output: %f\n", Float64(tc))
		fmt.Println("---")
	}
}

func TestInt(t *testing.T) {
	for _, tc := range testCases {
		fmt.Printf("Input:  %v\n", tc)
		fmt.Printf("Output: %d\n", Int(tc))
		fmt.Println("---")
	}
}

func BenchmarkConvert(b *testing.B) {
	b.Run("Int", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Int("-1")
		}
	})
	b.Run("Float64", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, tc := range testCases {
				Float64(tc)
			}
		}
	})
	b.Run("String", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, tc := range testCases {
				String(tc)
			}
		}
	})
	b.Run("Bool", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, tc := range testCases {
				Bool(tc)
			}
		}
	})
}
