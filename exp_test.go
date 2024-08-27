package main

import (
	"testing"
	"time"
)

func BenchmarkOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		time.Sleep(110 * time.Millisecond)
	}
}

func BenchmarkTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		time.Sleep(180 * time.Millisecond)
	}
}

func TestOne(t *testing.T) {
	t.Log("OK")
}

func TestTwo(t *testing.T) {
	t.Log("OK")
}
