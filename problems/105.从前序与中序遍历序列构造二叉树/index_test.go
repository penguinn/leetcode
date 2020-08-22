package index

import (
	"github.com/penguinn/leetcode/common"
	"reflect"
	"testing"
)

func Test_buildTree(t *testing.T) {
	type args struct {
		preorder []int
		inorder  []int
	}
	tests := []struct {
		name string
		args args
		want *common.TreeNode
	}{
		{
			name: "test1",
			args: args{preorder: []int{3, 9, 20, 15, 7}, inorder: []int{9, 3, 15, 20, 7}},
			want: common.NewBinaryTree([]interface{}{3, 9, 20, nil, nil, 15, 7}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildTree(tt.args.preorder, tt.args.inorder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
