package index

import "testing"

func Test_isOneBitCharacter(t *testing.T) {
	type args struct {
		bits []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{bits: []int{1, 0, 0}},
			want: true,
		},
		{
			name: "test2",
			args: args{bits: []int{1, 1, 1, 0}},
			want: false,
		},
		{
			name: "test3",
			args: args{bits: []int{1, 1, 1, 1, 0}},
			want: true,
		},
		{
			name: "test4",
			args: args{bits: []int{1, 0}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isOneBitCharacter(tt.args.bits); got != tt.want {
				t.Errorf("isOneBitCharacter() = %v, want %v", got, tt.want)
			}
		})
	}
}
