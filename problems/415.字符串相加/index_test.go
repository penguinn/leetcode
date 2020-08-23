package index

import "testing"

func Test_addStrings(t *testing.T) {
	type args struct {
		num1 string
		num2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{
				num1: "123",
				num2: "123",
			},
			want: "246",
		},
		{
			name: "test1",
			args: args{
				num1: "999",
				num2: "1",
			},
			want: "1000",
		},
		{
			name: "test1",
			args: args{
				num1: "1",
				num2: "999",
			},
			want: "1000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addStrings(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("addStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
