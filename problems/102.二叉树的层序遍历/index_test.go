package index

import (
	"reflect"
	"testing"

	"github.com/penguinn/leetcode/common"
)

func Test_levelOrder(t *testing.T) {
	type args struct {
		root *common.TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "test1",
			args: args{root: common.NewBinaryTree([]interface{}{3, 9, 20, nil, nil, 15, 7})},
			want: [][]int{{3}, {9, 20}, {15, 7}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
