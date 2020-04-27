package index

import (
	"reflect"
	"testing"

	"github.com/penguinn/leetcode/common"
)

// 二分法
// 从中序遍历变成了前序遍历
func Test_sortedArrayToBST(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want *common.TreeNode
	}{
		{
			name: "test1",
			args: args{nums: []int{-10, -3, 0, 5, 9}},
			want: common.NewBinaryTree([]interface{}{0, -3, 9, -10, nil, 5}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedArrayToBST(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedArrayToBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
