package heapsort

import (
	"reflect"
	"testing"
)

func Test_heapsort(t *testing.T) {
	l := []int{50, 30, 20, 60, 10, 8, 16, 15}
	want := []int{8, 10, 15, 16, 20, 30, 50, 60}

	heapsort(l)

	if !reflect.DeepEqual(want, l) {
		t.Errorf("heap sort failed.\nwant: %+v\ngot:  %+v", want, l)
	}
}

func Test_heapify(t *testing.T) {
	l := []int{50, 30, 20, 60, 10, 8, 16, 15}
	want := []int{60, 50, 20, 30, 10, 8, 16, 15}

	heapify(l, 0, len(l)-1)

	if !reflect.DeepEqual(want, l) {
		t.Errorf("heapify failed.\nwant: %+v\ngot:  %+v", want, l)
	}
}
