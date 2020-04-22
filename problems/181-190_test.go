package problems

import "testing"

func Test_maxProfit4(t *testing.T) {
	type args struct {
		k      int
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "测试1",
			args: args{
				k:      2,
				prices: []int{7, 1, 5, 3, 6, 4},
			},
			want: 7,
		},
		{
			name: "测试2",
			args: args{
				k:      3,
				prices: []int{8, 6, 2, 7, 6, 8, 1, 9},
			},
			want: 15,
		},
		{
			name: "测试3",
			args: args{
				prices: []int{7, 6, 4, 3, 1},
			},
			want: 0,
		},
		{
			name: "测试4",
			args: args{
				prices: []int{3, 3},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit4(tt.args.k, tt.args.prices); got != tt.want {
				t.Errorf("maxProfit4() = %v, want %v", got, tt.want)
			}
		})
	}
}
