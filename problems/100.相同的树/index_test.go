package index

import (
	"testing"

	"github.com/penguinn/leetcode/common"
)

func Test_isSameTree(t *testing.T) {
	type args struct {
		p *common.TreeNode
		q *common.TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{common.NewBinaryTree([]interface{}{1, 2, 3}), common.NewBinaryTree([]interface{}{1, 2, 3})},
			want: true,
		},
		{
			name: "test2",
			args: args{common.NewBinaryTree([]interface{}{1, 2}), common.NewBinaryTree([]interface{}{1, nil, 2})},
			want: false,
		},
		{
			name: "test3",
			args: args{common.NewBinaryTree([]interface{}{1, 2, 1}), common.NewBinaryTree([]interface{}{1, 1, 2})},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSameTree(tt.args.p, tt.args.q); got != tt.want {
				t.Errorf("isSameTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
