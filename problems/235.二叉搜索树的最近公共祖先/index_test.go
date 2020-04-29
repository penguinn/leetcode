package index

import (
	"reflect"
	"testing"

	"github.com/penguinn/leetcode/common"
)

func Test_lowestCommonAncestor(t *testing.T) {
	type args struct {
		root *common.TreeNode
		p    *common.TreeNode
		q    *common.TreeNode
	}
	tests := []struct {
		name string
		args args
		want *common.TreeNode
	}{
		{
			name: "test1",
			args: args{
				root: common.NewBinaryTree([]interface{}{6, 2, 8, 0, 4, 7, 9, nil, nil, 3, 5}),
				p:    &common.TreeNode{Val: 2},
				q:    &common.TreeNode{Val: 8},
			},
			want: &common.TreeNode{Val: 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lowestCommonAncestor(tt.args.root, tt.args.p, tt.args.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lowestCommonAncestor() = %v, want %v", got, tt.want)
			}
		})
	}
}
