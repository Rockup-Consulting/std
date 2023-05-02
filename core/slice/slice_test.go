package slice_test

import (
	"reflect"
	"testing"

	"github.com/Rockup-Consulting/go_std/core/slice"
)

var (
	testSlice = []int{1, 2, 2, 4, 5, 6, 6}
)

func TestContains(t *testing.T) {
	ok := slice.Contains(testSlice, 2)

	if !ok {
		t.Errorf("slice.Contains() = %t; want = %t", ok, true)
	}

	ok = slice.Contains(testSlice, 44)

	if ok {
		t.Errorf("slice.Contains() = %t; want = %t", ok, true)
	}
}

func TestMax(t *testing.T) {
	max := slice.Max(testSlice)
	if max != 6 {
		t.Errorf("slice.Max() = %d; want = %d", max, 6)
	}

}

func TestMin(t *testing.T) {
	min := slice.Min(testSlice)
	if min != 1 {
		t.Errorf("slice.Max() = %d; want = %d", min, 1)
	}
}

func TestRemoveDuplicates(t *testing.T) {
	want := []int{1, 2, 4, 5, 6}
	got := slice.RemoveDuplicates(testSlice)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("slice.RemoveDuplicates() = %v; want = %v", got, want)
	}
}

func TestPrepend(t *testing.T) {
	want := []int{1, 2, 4, 5, 6}
	x := []int{2, 4, 5, 6}

	got := slice.Prepend(1, x)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("slice.Prepend() = %v; want = %v", got, want)
	}
}
