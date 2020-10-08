package index

import (
	"reflect"
	"testing"

	. "github.com/penguinn/leetcode/common"
)

func Test_reverseBetween(t *testing.T) {
	type args struct {
		head *ListNode
		m    int
		n    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "test1",
			args: args{head: NewList([]interface{}{1, 2, 3, 4, 5}), m: 2, n: 4},
			want: NewList([]interface{}{1, 4, 3, 2, 5}),
		},
		{
			name: "test2",
			args: args{head: NewList([]interface{}{3,5}), m: 1, n: 2},
			want: NewList([]interface{}{5,3}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseBetween(tt.args.head, tt.args.m, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}
