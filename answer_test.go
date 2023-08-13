package answer

import (
	"reflect"
	"testing"
)

func Test_addNumbers(t *testing.T) {
	if !reflect.DeepEqual(divisor(6), []int{1, 2, 3, 6}) {
		t.Error("実際の結果が異なります。得られた結果＝＞", divisor(6))
	}

	if !reflect.DeepEqual(divisor(7), []int{1, 7}) {
		t.Error("実際の結果が異なります。得られた結果＝＞", divisor(7))
	}
}

func Test_q1(t *testing.T) {
	if !reflect.DeepEqual(q1(583305), []int{1, 3, 5, 15, 37, 111, 185, 555, 1051, 3153, 5255, 15765, 38887, 116661, 194435, 583305}) {
		t.Error("実際の結果が異なります。得られた結果＝＞", q1(583305))
	}
}
