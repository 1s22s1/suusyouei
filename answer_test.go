package answer

import (
	"reflect"
	"testing"
)

func TestDivisor(t *testing.T) {
	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "normal1",
			args: args{n: 6},
			want: []int{1, 2, 3, 6},
		},
		{
			name: "normal2",
			args: args{n: 7},
			want: []int{1, 7},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divisor(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("divisor() + %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsAbundant(t *testing.T) {
	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "normal1",
			args: args{n: 6},
			want: false,
		},
		{
			name: "normal2",
			args: args{n: 12},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAbundant(tt.args.n); got != tt.want {
				t.Errorf("isAbundant() + %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCollatz(t *testing.T) {
	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "normal1",
			args: args{n: 1},
			want: []int{1, 4, 2, 1},
		},
		{
			name: "normal2",
			args: args{n: 3},
			want: []int{3, 10, 5, 16, 8, 4, 2, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := collatz(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("collatz() + %v, want %v", got, tt.want)
			}
		})
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

func TestQ2(t *testing.T) {
	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "normal1",
			args: args{n: 100000},
			want: 100000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := q2(tt.args.n); got != tt.want {
				t.Errorf("q2() + %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQ3(t *testing.T) {
	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "normal1",
			args: args{n: 1},
			want: 20830,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := q3(tt.args.n); got != tt.want {
				t.Errorf("q3() + %v, want %v", got, tt.want)
			}
		})
	}
}
