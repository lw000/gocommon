package tyutils

import "testing"

func TestReverse(t *testing.T) {
	type args struct {
		src string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Reverse", args: args{src: "1234"}, want: "4321"},
		{name: "Reverse", args: args{src: "abcd"}, want: "dcba"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse(tt.args.src); got != tt.want {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
