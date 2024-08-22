package leetcode

import "testing"

func TestNumberSquares_Calc(t *testing.T) {
	type args struct {
		number int
	}
	tests := []struct {
		name string
		ns   *NumberSquares
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			"case1",
			&NumberSquares{},
			args{12},
			3,
		},
		{
			"case2",
			&NumberSquares{},
			args{13},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ns.Calc(tt.args.number); got != tt.want {
				t.Errorf("NumberSquares.Calc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumberSquares_min(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		ns   *NumberSquares
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ns.min(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("NumberSquares.min() = %v, want %v", got, tt.want)
			}
		})
	}
}
