package index

import (
	"testing"

	"github.com/penguinn/leetcode/common"
)

func Test_isSubtree(t *testing.T) {
	type args struct {
		s *common.TreeNode
		t *common.TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{
				s: common.NewBinaryTree([]interface{}{3, 4, 5, 1, 2}),
				t: common.NewBinaryTree([]interface{}{4, 1, 2}),
			},
			want: true,
		},
		{
			name: "test2",
			args: args{
				s: common.NewBinaryTree([]interface{}{3, 4, 5, 1, nil, 2}),
				t: common.NewBinaryTree([]interface{}{3, 1, 2}),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubtree(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isSubtree() = %v, want %v", got, tt.want)
			}
		})
	}
}
