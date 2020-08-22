package index

import (
	"reflect"
	"testing"

	"github.com/penguinn/leetcode/common"
)

func Test_addTwoNumbers(t *testing.T) {
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
				l1: common.NewList([]interface{}{7, 2, 4, 3}),
				l2: common.NewList([]interface{}{5, 6, 4}),
			},
			want: common.NewList([]interface{}{7, 8, 0, 7}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
