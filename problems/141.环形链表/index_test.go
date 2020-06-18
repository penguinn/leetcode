package index

import (
	"testing"

	"github.com/penguinn/leetcode/common"
)

func Test_hasCycle(t *testing.T) {
	type args struct {
		head *common.ListNode
	}
	list := common.NewList([]interface{}{3, 2, 0, -4})
	list.Next.Next.Next.Next = list.Next

	list2 := common.NewList([]interface{}{1, 2})
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{head: list},
			want: true,
		},
		{
			name: "test2",
			args: args{head: list2},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycle(tt.args.head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
