package index

import "testing"

func Test_translateNum(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test0",
			args: args{num: 2},
			want: 1,
		},
		{
			name: "test1",
			args: args{num: 25},
			want: 2,
		},
		{
			name: "test2",
			args: args{num: 12258},
			want: 5,
		},
		{
			name: "test3",
			args: args{num: 1068385902},
			want: 2,
		},
		{
			name: "test1",
			args: args{num: 11111111111},
			want: 144,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := translateNum(tt.args.num); got != tt.want {
				t.Errorf("translateNum() = %v, want %v", got, tt.want)
			}
		})
	}
}
