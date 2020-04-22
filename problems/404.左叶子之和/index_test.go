package index

import (
	"github.com/penguinn/leetcode/common"
	"testing"
)

func Test_sumOfLeftLeaves(t *testing.T) {
	type args struct {
		root *common.TreeNode
	}
	root := common.NewBinaryTree([]interface{}{3, 9, 20, nil, nil, 15, 7})
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "测试1",
			args: args{root: root},
			want: 24,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumOfLeftLeaves(tt.args.root); got != tt.want {
				t.Errorf("sumOfLeftLeaves() = %v, want %v", got, tt.want)
			}
		})
	}
}
