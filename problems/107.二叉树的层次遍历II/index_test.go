package index

import (
	"reflect"
	"testing"

	"github.com/penguinn/leetcode/common"
)

func Test_levelOrderBottom(t *testing.T) {
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
			args: args{common.NewBinaryTree([]interface{}{3, 9, 20, nil, nil, 15, 7})},
			want: [][]int{{15, 7}, {9, 20}, {3}},
		},
		{
			name: "test2",
			args: args{common.NewBinaryTree([]interface{}{1,2})},
			want: [][]int{{2}, {1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrderBottom(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrderBottom() = %v, want %v", got, tt.want)
			}
		})
	}
}
