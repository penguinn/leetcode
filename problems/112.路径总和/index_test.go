package index

import (
	"testing"

	"github.com/penguinn/leetcode/common"
)

func Test_hasPathSum(t *testing.T) {
	type args struct {
		root *common.TreeNode
		sum  int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{root: common.NewBinaryTree([]interface{}{5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, nil, nil, nil, 1})},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasPathSum(tt.args.root, tt.args.sum); got != tt.want {
				t.Errorf("hasPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
