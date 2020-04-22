package problems

import (
	"testing"
)

func Test_complexNumberMultiply(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "测试1",
			args: args{
				a: "1+1i",
				b: "1+1i",
			},
			want: "0+2i",
		},
		{
			name: "测试2",
			args: args{
				a: "1+-1i",
				b: "1+-1i",
			},
			want: "0+-2i",
		},
		{
			name: "测试3",
			args: args{
				a: "1+-1i",
				b: "0+0i",
			},
			want: "0+0i",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ComplexNumberMultiply(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("complexNumberMultiply() = %v, want %v", got, tt.want)
			}
		})
	}
}
