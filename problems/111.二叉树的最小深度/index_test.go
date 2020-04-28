package index

import (
	"testing"

	"github.com/penguinn/leetcode/common"
)

func Test_minDepth(t *testing.T) {
	type args struct {
		root *common.TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{root: common.NewBinaryTree([]interface{}{3, 9, 20, nil, nil, 15, 7})},
			want: 2,
		},
		{
			name: "test2",
			args: args{root: common.NewBinaryTree([]interface{}{1, 2})},
			want: 2,
		},
		{
			name: "test3",
			args: args{root: common.NewBinaryTree([]interface{}{1, 2, 3, 4, nil, nil, 5})},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDepth(tt.args.root); got != tt.want {
				t.Errorf("minDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
