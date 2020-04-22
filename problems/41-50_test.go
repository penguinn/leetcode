package problems

import (
	"testing"
)

func Test_jump(t *testing.T) {
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
			args: args{
				nums: []int{2, 3, 1, 1, 4},
			},
			want: 2,
		},
		{
			name: "测试2",
			args: args{
				nums: []int{1, 1},
			},
			want: 1,
		},
		{
			name: "测试3",
			args: args{
				nums: []int{1, 1, 1, 1},
			},
			want: 3,
		},
		{
			name: "测试4",
			args: args{
				nums: []int{5, 9, 3, 2, 1, 0, 2, 3, 3, 1, 0, 0},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := jump(tt.args.nums); got != tt.want {
				t.Errorf("jump() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multiply(t *testing.T) {
	type args struct {
		num1 string
		num2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "测试1",
			args: args{
				num1: "11",
				num2: "11",
			},
			want: "121",
		},
		{
			name: "测试1",
			args: args{
				num1: "9",
				num2: "99",
			},
			want: "891",
		},
		{
			name: "测试1",
			args: args{
				num1: "123",
				num2: "456",
			},
			want: "56088",
		},
		{
			name: "测试1",
			args: args{
				num1: "9133",
				num2: "0",
			},
			want: "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := multiply(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}
