package iplocation

import (
	"testing"
)

func init() {
	FillCache()
}

func BenchmarkMaxMind(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetLocationByIp("182.74.246.102")
//		PrintLocation(loc)
	}

}

func BenchmarkIpTree(b *testing.B)  {
	for n := 0; n < b.N; n++ {
		GetLocationByIpTree("182.74.246.102")
		//PrintLocation(loc)
	}
}
