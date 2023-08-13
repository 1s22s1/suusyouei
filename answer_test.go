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

func TestQ1(t *testing.T) {
	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "normal",
			args: args{n: 583305},
			want: []int{1, 3, 5, 15, 37, 111, 185, 555, 1051, 3153, 5255, 15765, 38887, 116661, 194435, 583305},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := q1(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("q1() + %v, want %v", got, tt.want)
			}
		})
	}
}
