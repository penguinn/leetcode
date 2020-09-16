package index

import "testing"

func Test_maxProduct(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{nums: []int{2, 3, -2, 4}},
			want: 6,
		},
		{
			name: "test2",
			args: args{nums: []int{-2, 0, -1}},
			want: 0,
		},
		{
			name: "test3",
			args: args{nums: []int{-2, 4, -1}},
			want: 8,
		},
		{
			name: "test4",
			args: args{nums: []int{2, -5, -2, -4, 3}},
			want: 24,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProduct(tt.args.nums); got != tt.want {
				t.Errorf("maxProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
