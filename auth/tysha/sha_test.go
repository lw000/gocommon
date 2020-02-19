package tysha

import "testing"

func TestSha1(t *testing.T) {
	type args struct {
		str []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "TestSha1", args: args{str: []byte("111111111111111111")}, want: "352716cecf34d8a4d1c2e570a05b5818ff857001"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sha1(tt.args.str); got != tt.want {
				t.Errorf("Sha1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSha224(t *testing.T) {
	type args struct {
		str []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "TestSha1", args: args{str: []byte("111111111111111111")}, want: "c2444859dcf2c1e39babc58933d87fff0561e980dd45b746d984b682"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sha224(tt.args.str); got != tt.want {
				t.Errorf("Sha224() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSha256(t *testing.T) {
	type args struct {
		str []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "TestSha1", args: args{str: []byte("111111111111111111")}, want: "26980a9bf7dba794bd14eae90dbda34109d37f66950c16cce7b87dffd4535b40"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sha256(tt.args.str); got != tt.want {
				t.Errorf("Sha256() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSha512(t *testing.T) {
	type args struct {
		str []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "TestSha1", args: args{str: []byte("111111111111111111")}, want: "87ba0d56d929d2e3013ba020592e50c79d70c5ec75e73bb76373f0ad44517ef2c3c5901b17cf23df49a226ac6f52f672130025e1b15d9d9a99236bffd0c84ca5"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sha512(tt.args.str); got != tt.want {
				t.Errorf("Sha512() = %v, want %v", got, tt.want)
			}
		})
	}
}
