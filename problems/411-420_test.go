package problems

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
			name: "测试1",
			args: args{
				num1: "443",
				num2: "777",
			},
			want: "1220",
		},
		{
			name: "测试1",
			args: args{
				num1: "0",
				num2: "1",
			},
			want: "1",
		},
		{
			name: "测试1",
			args: args{
				num1: "123",
				num2: "123",
			},
			want: "246",
		},
		{
			name: "测试1",
			args: args{
				num1: "9",
				num2: "99",
			},
			want: "108",
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
