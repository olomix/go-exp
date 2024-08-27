package main

import (
	"testing"
	"time"
)

func BenchmarkOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		time.Sleep(90 * time.Millisecond)
	}
}

func BenchmarkTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		time.Sleep(190 * time.Millisecond)
	}
}

func TestOne(t *testing.T) {
	t.Log("OK")
}

func TestTwo(t *testing.T) {
	t.Log("OK")
}
