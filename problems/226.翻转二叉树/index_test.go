package index

import (
	"reflect"
	"testing"

	"github.com/penguinn/leetcode/common"
)

func Test_invertTree(t *testing.T) {
	type args struct {
		root *common.TreeNode
	}
	tests := []struct {
		name string
		args args
		want *common.TreeNode
	}{
		{
			name: "test1",
			args: args{root: common.NewBinaryTree([]interface{}{4, 2, 7, 1, 3, 6, 9})},
			want: common.NewBinaryTree([]interface{}{4, 7, 2, 9, 6, 3, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := invertTree(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("invertTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
