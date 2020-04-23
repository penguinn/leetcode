package index

import (
	"reflect"
	"testing"

	"github.com/penguinn/leetcode/common"
)

func Test_findMode(t *testing.T) {
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
			args: args{root: common.NewBinaryTree([]interface{}{1, nil, 2, nil, nil, 2})},
			want: []int{2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMode(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findMode() = %v, want %v", got, tt.want)
			}
		})
	}
}
