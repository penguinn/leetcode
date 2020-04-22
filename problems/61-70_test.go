package problems

import (
	"testing"
)

func TestUniquePaths(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "测试1",
			args: args{
				m: 2,
				n: 2,
			},
			want: 2,
		},
		{
			name: "测试2",
			args: args{
				m: 3,
				n: 3,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePaths(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("UniquePaths() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_uniquePathsWithObstacles(t *testing.T) {
	type args struct {
		obstacleGrid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "测试1",
			args: args{
				obstacleGrid: [][]int{
					[]int{0, 0, 0},
					[]int{0, 1, 0},
					[]int{0, 0, 0},
				},
			},
			want: 2,
		},
		{
			name: "测试2",
			args: args{
				obstacleGrid: [][]int{
					[]int{0},
				},
			},
			want: 1,
		},
		{
			name: "测试3",
			args: args{
				obstacleGrid: [][]int{
					[]int{1},
				},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePathsWithObstacles(tt.args.obstacleGrid); got != tt.want {
				t.Errorf("uniquePathsWithObstacles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minPathSum(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "测试1",
			args: args{
				grid: [][]int{
					{1, 3, 1},
					{1, 5, 1},
					{4, 2, 1},
				},
			},
			want: 7,
		},
		{
			name: "测试2",
			args: args{
				grid: [][]int{
					{1},
				},
			},
			want: 1,
		},
		{
			name: "测试3",
			args: args{
				grid: [][]int{},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minPathSum(tt.args.grid); got != tt.want {
				t.Errorf("minPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
