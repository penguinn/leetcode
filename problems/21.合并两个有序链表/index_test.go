package index

import (
	"reflect"
	"testing"

	"github.com/penguinn/leetcode/common"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		l1 *common.ListNode
		l2 *common.ListNode
	}
	tests := []struct {
		name string
		args args
		want *common.ListNode
	}{
		{
			name: "test1",
			args: args{
				l1: common.NewList([]interface{}{1, 2, 4}),
				l2: common.NewList([]interface{}{1, 3, 4}),
			},
			want: common.NewList([]interface{}{1, 1, 2, 3, 4, 4}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
