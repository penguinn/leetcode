package problems

import "testing"

func Test_orderOfLargestPlusSign(t *testing.T) {
	type args struct {
		N     int
		mines [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "测试1",
			args: args{
				N:     5,
				mines: [][]int{{4, 2}},
			},
			want: 2,
		},
		{
			name: "测试2",
			args: args{
				N:     2,
				mines: [][]int{},
			},
			want: 1,
		},
		{
			name: "测试3",
			args: args{
				N:     1,
				mines: [][]int{{0, 0}},
			},
			want: 0,
		},
		{
			name: "测试4",
			args: args{
				N:     3,
				mines: [][]int{{0, 1}},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := orderOfLargestPlusSign(tt.args.N, tt.args.mines); got != tt.want {
				t.Errorf("orderOfLargestPlusSign() = %v, want %v", got, tt.want)
			}
		})
	}
}
