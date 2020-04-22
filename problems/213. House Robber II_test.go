package problems

import "testing"

func Test_rob(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "测试1",
			args: args{nums: []int{2, 3, 2}},
			want: 3,
		},
		{
			name: "测试2",
			args: args{nums: []int{1, 2, 3, 1}},
			want: 4,
		},
		{
			name: "测试2",
			args: args{nums: []int{1, 7, 9, 2}},
			want: 10,
		},
		{
			name: "测试2",
			args: args{nums: []int{6, 6, 4, 8, 4, 3, 3, 10}},
			want: 27,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob(tt.args.nums); got != tt.want {
				t.Errorf("rob() = %v, want %v", got, tt.want)
			}
		})
	}
}
