package index

import (
	"reflect"
	"testing"
)

func Test_merge(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "test1",
			args: args{intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}},
			want: [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			name: "test1",
			args: args{intervals: [][]int{{15, 18}, {8, 10}, {1, 3}, {2, 6}}},
			want: [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			name: "test1",
			args: args{intervals: [][]int{{15, 18}, {3, 4}, {8, 10}, {1, 3}, {2, 6}}},
			want: [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			name: "test1",
			args: args{intervals: [][]int{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}}},
			want: [][]int{{1, 10}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge1(tt.args.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
