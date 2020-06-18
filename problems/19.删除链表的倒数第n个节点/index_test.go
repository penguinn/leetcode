package index

import (
	"reflect"
	"testing"

	"github.com/penguinn/leetcode/common"
)

func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		head *common.ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want *common.ListNode
	}{
		{
			name: "test1",
			args: args{head: common.NewList([]interface{}{1, 2, 3, 4, 5}), n: 2},
			want: common.NewList([]interface{}{1, 2, 3, 5}),
		},
		{
			name: "test2",
			args: args{head: common.NewList([]interface{}{1, 2}), n: 1},
			want: common.NewList([]interface{}{1}),
		},
		{
			name: "test3",
			args: args{head: common.NewList([]interface{}{1, 2}), n: 2},
			want: common.NewList([]interface{}{2}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromEnd(tt.args.head, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
