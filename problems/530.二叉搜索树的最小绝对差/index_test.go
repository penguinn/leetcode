package index

import (
	"testing"

	"github.com/penguinn/leetcode/common"
)

func Test_getMinimumDifference(t *testing.T) {
	type args struct {
		root *common.TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test",
			args: args{root: common.NewBinaryTree([]interface{}{1, nil, 3, nil, nil, 2})},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMinimumDifference(tt.args.root); got != tt.want {
				t.Errorf("getMinimumDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
