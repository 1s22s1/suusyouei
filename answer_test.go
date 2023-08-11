package answer

import "testing"

func Test_addNumbers(t *testing.T) {
	result := addNumbers(2, 3)

	if result != 5 {
		t.Error("実際の結果が異なります。得られた結果＝＞", result)
	}
}
