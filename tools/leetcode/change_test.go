package leetcode

import "testing"

func TestChange_Calc(t *testing.T) {
	type args struct {
		amount int
		coins  []int
	}
	tests := []struct {
		name string
		c    *Change
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			"case1",
			&Change{},
			args{
				5,
				[]int{1, 2, 5},
			},
			4,
		},
		{
			"case1",
			&Change{},
			args{
				3,
				[]int{2},
			},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Calc(tt.args.amount, tt.args.coins); got != tt.want {
				t.Errorf("Change.Calc() = %v, want %v", got, tt.want)
			}
		})
	}
}
