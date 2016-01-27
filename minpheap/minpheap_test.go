package minpheap

import (
	"math/rand"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMinPHeap(t *testing.T) {
	Convey("Test that output after random input is sorted.", t, func() {
		m := New()
		for i := 0; i < 1000; i++ {
			m.Insert(i, rand.Float32())
		}
		_, prev := m.Pop()
		for i := 0; i < 999; i++ {
			_, key := m.Pop()
			So(key, ShouldBeGreaterThanOrEqualTo, prev)
			prev = key
		}
	})

	Convey("Test DecreaseKey.", t, func() {
		m := New()
		for i := 0; i < 1000; i++ {
			m.Insert(i, rand.Float32())
		}
		for i := 0; i < 1000; i++ {
			val, _ := m.PeekAtVal(i)
			m.DecreaseKey(i, val-1)
		}
		_, prev := m.Pop()
		for i := 0; i < 999; i++ {
			_, key := m.Pop()
			So(key, ShouldBeGreaterThanOrEqualTo, prev)
			prev = key
		}
	})
}

func BenchmarkInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		nums := make([]float32, 1000)
		for j := 0; j < 1000; j++ {
			nums[j] = rand.Float32()
		}
		m := New()
		b.StartTimer()
		for j := 0; j < 1000; j++ {
			m.Insert("", nums[j])
		}
	}
}

func BenchmarkPop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		nums := make([]float32, 1000)
		for j := 0; j < 1000; j++ {
			nums[j] = rand.Float32()
		}
		m := New()
		for j := 0; j < 1000; j++ {
			m.Insert("", nums[j])
		}
		b.StartTimer()
		for j := 0; j < 1000; j++ {
			m.Pop()
		}
	}
}
