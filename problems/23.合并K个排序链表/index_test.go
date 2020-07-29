package index

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/penguinn/leetcode/common"
)

func Test_mergeKLists(t *testing.T) {
	type args struct {
		lists []*common.ListNode
	}
	lists := []*common.ListNode{}
	lists = append(lists, common.NewList([]interface{}{1, 4, 5}))
	lists = append(lists, common.NewList([]interface{}{1, 3, 4}))
	lists = append(lists, common.NewList([]interface{}{2, 6}))
	tests := []struct {
		name string
		args args
		want *common.ListNode
	}{
		{
			name: "test1",
			args: args{
				lists: lists,
			},
			want: common.NewList([]interface{}{1, 1, 2, 3, 4, 4, 5, 6}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeKLists(tt.args.lists); !reflect.DeepEqual(got, tt.want) {
				for got != nil {
					fmt.Println(got.Val)
					got = got.Next
				}
				t.Errorf("mergeKLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
