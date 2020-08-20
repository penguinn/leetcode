package index

import (
	"reflect"
	"testing"

	"github.com/penguinn/leetcode/common"
)

func Test_rightSideView(t *testing.T) {
	type args struct {
		root *common.TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{root: common.NewBinaryTree([]interface{}{1, 2, 3, nil, 5, nil, 4})},
			want: []int{1, 3, 4},
		},
		{
			name: "test1",
			args: args{root: common.NewBinaryTree([]interface{}{1, 2, 3, nil, 5, nil, 4, nil, nil, 9})},
			want: []int{1, 3, 4, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rightSideView(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rightSideView() = %v, want %v", got, tt.want)
			}
		})
	}
}
