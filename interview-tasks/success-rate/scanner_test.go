package main

import (
	"testing"
	"time"
)

func BenchmarkScanner_CalcAppsRatesWithChan(b *testing.B) {
	s := NewScanner(time.Second * 2)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		s.CalcAppsRatesWithChan(100)
	}
}

func BenchmarkScanner_CalcAppsRatesWithMux(b *testing.B) {
	s := NewScanner(time.Second * 2)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		s.CalcAppsRatesWithMux(100)
	}
}
