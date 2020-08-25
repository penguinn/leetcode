package index

import (
	"reflect"
	"testing"

	"github.com/penguinn/leetcode/common"
)

func Test_getIntersectionNode(t *testing.T) {
	type args struct {
		headA *common.ListNode
		headB *common.ListNode
	}
	tests := []struct {
		name string
		args args
		want *common.ListNode
	}{
		//{
		//	name: "test1",
		//	args: args{headA: common.NewList([]interface{}{4, 1, 8, 4, 5}), headB: common.NewList([]interface{}{5, 0, 1, 8, 4, 5})},
		//	want: common.NewList([]interface{}{1, 8, 4, 5}),
		//},
		//{
		//	name: "test2",
		//	args: args{headA: common.NewList([]interface{}{0, 9, 1, 2, 4}), headB: common.NewList([]interface{}{3, 2, 4})},
		//	want: common.NewList([]interface{}{2, 4}),
		//},
		{
			name: "test3",
			args: args{headA: common.NewList([]interface{}{2, 6, 4}), headB: common.NewList([]interface{}{1, 5})},
			want: nil,
		},
		{
			name: "test3",
			args: args{headA: common.NewList([]interface{}{2, 6}), headB: common.NewList([]interface{}{1, 5})},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getIntersectionNode(tt.args.headA, tt.args.headB); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getIntersectionNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
