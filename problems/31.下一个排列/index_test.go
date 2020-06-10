package index

import "testing"

func Test_nextPermutation(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test1",
			args: args{nums: []int{1, 2, 3}},
		},
		{
			name: "test1",
			args: args{nums: []int{3, 2, 1}},
		},
		{
			name: "test1",
			args: args{nums: []int{1, 1, 5}},
		},
		{
			name: "test1",
			args: args{nums: []int{1, 3, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nextPermutation(tt.args.nums)
		})
	}
}
