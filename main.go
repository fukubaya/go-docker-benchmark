package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func main() {
	result := testing.Benchmark(func(b *testing.B) {
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			l := []int{}
			for j := 0; j < 100000; j++ {
				l = append(l, j)
			}
			json.Marshal(l)
		}
	})
	fmt.Printf("%s\n", result)
}
