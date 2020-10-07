package index

import (
	"reflect"
	"testing"

	. "github.com/penguinn/leetcode/common"
)

func Test_reverseKGroup(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "test1",
			args: args{
				head: NewList([]interface{}{1, 2, 3, 4, 5}),
				k:    2,
			},
			want: NewList([]interface{}{2, 1, 4, 3, 5}),
		},
		{
			name: "test2",
			args: args{
				head: NewList([]interface{}{1, 2, 3, 4, 5}),
				k:    3,
			},
			want: NewList([]interface{}{3, 2, 1, 4, 5}),
		},
		{
			name: "test3",
			args: args{
				head: NewList([]interface{}{1, 2, 3, 4, 5}),
				k:    1,
			},
			want: NewList([]interface{}{1, 2, 3, 4, 5}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseKGroup(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseKGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}
