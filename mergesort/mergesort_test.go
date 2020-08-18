package mergesort

import (
	"reflect"
	"testing"
)

func Test_mergesort(t *testing.T) {
	l := []int{9, 3, 7, 5, 6, 4, 8, 2}
	got := mergeSort(l)
	want := []int{2, 3, 4, 5, 6, 7, 8, 9}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("mergesort failed.\nwant: %+v\ngot:  %+v", want, got)
	}
}
