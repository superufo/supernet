package utils

import (
	"log"
	"testing"
)

// TestGenRandString xxx
func TestGenRandString(t *testing.T) {
	n := 10
	c := 10000
	m := make(map[string]int)
	for i := 0; i < c; i++ {
		r := GenRandString(n)
		m[r] = i
		log.Println(r)
	}
	log.Println(len(m) == c)
}

// BenchmarkRandString bench rand
func BenchmarkRandString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		go GenRandString(20)
	}
}
