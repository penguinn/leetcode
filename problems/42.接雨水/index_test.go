package index

import "testing"

func Test_trap(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}},
			want: 6,
		},
		{
			name: "test2",
			args: args{height: []int{2, 0, 2}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trap2(tt.args.height); got != tt.want {
				t.Errorf("trap() = %v, want %v", got, tt.want)
			}
		})
	}
}
