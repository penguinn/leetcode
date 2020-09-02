package index

import (
	"testing"

	"github.com/penguinn/leetcode/common"
)

func Test_reorderList(t *testing.T) {
	type args struct {
		head *common.ListNode
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test1",
			args: args{head: common.NewList([]interface{}{1, 2, 3, 4})},
		},
		{
			name: "test2",
			args: args{head: common.NewList([]interface{}{1, 2, 3, 4, 5})},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reorderList(tt.args.head)
		})
	}
}
