package gomapreduce

import (
	"reflect"
	"strconv"
	"testing"
)

func TestMap(t *testing.T) {
	ints_to_five := []int{1, 2, 3, 4, 5}
	added_strings := Map(func(orig int, i int) string { return strconv.Itoa(orig + i) }, ints_to_five)
	if !reflect.DeepEqual(added_strings, []string{"1", "3", "5", "7", "9"}) {
		t.Fail()
	}
}
