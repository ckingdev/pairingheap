package minpheap

import (
	"math/rand"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMinPHeap(t *testing.T) {
	Convey("Test that output after random input is sorted.", t, func() {
		m := New()
		t.Logf("Inserting...")
		for i := 0; i < 1000; i++ {
			m.Insert(0, rand.Float32())
		}
		_, prev := m.Pop()
		for i := 0; i < 999; i++ {
			_, key := m.Pop()
			So(key, ShouldBeGreaterThanOrEqualTo, prev)
			prev = key
		}
	})
}
