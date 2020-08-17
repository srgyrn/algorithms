package quicksort

import (
	"reflect"
	"testing"
)

func Test_quicksort(t *testing.T) {
	l := []int{10, 16, 8, 12, 15, 6, 3, 9, 5}
	quicksort(l, 0, len(l)-1)
	want := []int{3, 5, 6, 8, 9, 10, 12, 15, 16}

	if !reflect.DeepEqual(want, l) {
		t.Errorf("quicksort failed.\nwant: %+v\ngot:  %+v", want, l)
	}
}