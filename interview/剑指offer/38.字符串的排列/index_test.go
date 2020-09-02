package index

import (
	"reflect"
	"testing"
)

func Test_permutation(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test1",
			args: args{s: "abc"},
			want: []string{"abc", "acb", "bac", "bca", "cab", "cba"},
		},
		{
			name: "test1",
			args: args{s: "aab"},
			want: []string{"aba", "aab", "baa"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := permutation(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("permutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
