package index

import (
	"testing"

	"github.com/penguinn/leetcode/common"
)

func Test_pathSum(t *testing.T) {
	type args struct {
		root *common.TreeNode
		sum  int
	}
	root1 := common.NewBinaryTree([]interface{}{10, 5, -3, 3, 2, nil, 11, 3, -2, nil, 1})
	root2 := common.NewBinaryTree([]interface{}{1, nil, 2, nil, nil, nil, 3, nil, nil, nil, nil, nil, nil, nil, 4, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, 5})
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				root: root1,
				sum:  8,
			},
			want: 3,
		},
		{
			name: "test2",
			args: args{
				root: root2,
				sum:  3,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pathSum(tt.args.root, tt.args.sum); got != tt.want {
				t.Errorf("pathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
