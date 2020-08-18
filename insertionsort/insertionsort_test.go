package insertionsort

import (
	"reflect"
	"testing"
)

func Test_insertionSort(t *testing.T) {
	l := []int{6, 7, 3, 8, 1, 5, 9, 2, 4}
	want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	insertionSort(l)

	if !reflect.DeepEqual(l, want) {
		t.Errorf("insertion sort failed.\nwant: %+v\ngot:  %+v", want, l)
	}
}
