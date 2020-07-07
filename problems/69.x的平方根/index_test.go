package index

import (
	"testing"
)

func Test_mySqrt(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{x: 4},
			want: 2,
		},
		{
			name: "test2",
			args: args{x: 8},
			want: 2,
		},
		{
			name: "test3",
			args: args{x: 6},
			want: 2,
		},
		{
			name: "test4",
			args: args{x: 0},
			want: 0,
		},
		{
			name: "test5",
			args: args{x: 1},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mySqrt(tt.args.x); got != tt.want {
				t.Errorf("mySqrt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mySqrt6(t *testing.T) {
	type args struct {
		x float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "test1",
			args: args{x: 27},
			want: 5.196152,
		},
		{
			name: "test2",
			args: args{x: 35},
			want: 5.916080,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mySqrt6(tt.args.x); got != tt.want {
				t.Errorf("mySqrt6() = %v, want %v", got, tt.want)
			}
		})
	}
}
