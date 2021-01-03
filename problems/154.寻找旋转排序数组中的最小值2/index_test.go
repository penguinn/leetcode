package index

import "testing"

func Test_findMin(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name:"test1",
			args:args{nums:[]int{3,4,5,1,2}},
			want:1,
		},
		{
			name:"test2",
			args:args{nums:[]int{4,5,6,7,0,1,2}},
			want:0,
		},
		{
			name:"test3",
			args:args{nums:[]int{1}},
			want:1,
		},
		{
			name:"test4",
			args:args{nums:[]int{1,2,3,4,5,6}},
			want:1,
		},
		{
			name:"test5",
			args:args{nums:[]int{2,2,2,0,1}},
			want:0,
		},
		{
			name:"test6",
			args:args{nums:[]int{3,1,3}},
			want:1,
		},
		{
			name:"test7",
			args:args{nums:[]int{10,1,10,10,10}},
			want:1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMin(tt.args.nums); got != tt.want {
				t.Errorf("findMin() = %v, want %v", got, tt.want)
			}
		})
	}
}
