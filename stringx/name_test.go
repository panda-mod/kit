package stringx

import (
	"fmt"
	"testing"
)

var testCases = []string{
	"thequickbrownfox",
	"theQuickBrownFox",
	"TheQuickBrownFox",
	"The-Quick-Brown-Fox",
	"This is a \"quoted\" word",
	"Enable 2FA authentication",
	"MyURLParser",
	"iOS",
	"iPhone",
	"ABC123xyz",
	"TCP/IP",
	"123ABC",
	"CamelCase_UPPER_snake",
}

func TestWords(t *testing.T) {
	for _, tc := range testCases {
		fmt.Printf("Input:  %q\n", tc)
		fmt.Printf("Output: %q\n", Words(tc))
		fmt.Println("---")
	}
}

func TestKebabCase(t *testing.T) {
	for _, tc := range testCases {
		result := KebabCase(tc)
		fmt.Printf("Input:  %q\n", tc)
		fmt.Printf("Output: %q\n", result)
		fmt.Println("---")
	}
}

func TestSnakeCase(t *testing.T) {
	for _, tc := range testCases {
		result := SnakeCase(tc)
		fmt.Printf("Input:  %q\n", tc)
		fmt.Printf("Output: %q\n", result)
		fmt.Println("---")
	}
}

func Benchmark(b *testing.B) {
	b.Run("KebabCase", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, tc := range testCases {
				KebabCase(tc)
			}
		}
	})
	b.Run("SnakeCase", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, tc := range testCases {
				SnakeCase(tc)
			}
		}
	})
}
