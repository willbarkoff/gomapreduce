package gomapreduce_test

import (
	"strconv"
	"testing"

	"github.com/willbarkoff/gomapreduce"
)

func TestEqual(t *testing.T) {
	ints_to_five := []int{1, 2, 3, 4, 5}
	ints_to_four := []int{1, 2, 3, 4}
	if !gomapreduce.Equal(ints_to_five, ints_to_five) {
		t.Fail()
	}
	if gomapreduce.Equal(ints_to_four, ints_to_five) {
		t.Fail()
	}

}

func TestMap(t *testing.T) {
	ints_to_five := []int{1, 2, 3, 4, 5}
	added_strings := gomapreduce.Mapi(func(orig int, i int) string { return strconv.Itoa(orig + i) }, ints_to_five)
	if !gomapreduce.Equal(added_strings, []string{"1", "3", "5", "7", "9"}) {
		t.Fail()
	}
}
