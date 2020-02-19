package tyutils

import (
	"testing"
)

func TestMinInt64(t *testing.T) {
	type args struct {
		num1 int64
		num2 int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinInt64(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("MinInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMax(t *testing.T) {
	type args struct {
		num1 int
		num2 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Max(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxInt32(t *testing.T) {
	type args struct {
		num1 int32
		num2 int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxInt32(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("MaxInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxInt64(t *testing.T) {
	type args struct {
		num1 int64
		num2 int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxInt64(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("MaxInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMin(t *testing.T) {
	type args struct {
		num1 int
		num2 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Min(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("Min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinInt32(t *testing.T) {
	type args struct {
		num1 int32
		num2 int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinInt32(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("MinInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxmin(t *testing.T) {
	type args struct {
		num1 int
		num2 int
	}
	tests := []struct {
		name    string
		args    args
		wantMax int
		wantMin int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMax, gotMin := Maxmin(tt.args.num1, tt.args.num2)
			if gotMax != tt.wantMax {
				t.Errorf("Maxmin() gotMax = %v, want %v", gotMax, tt.wantMax)
			}
			if gotMin != tt.wantMin {
				t.Errorf("Maxmin() gotMin = %v, want %v", gotMin, tt.wantMin)
			}
		})
	}
}

func TestMaxminInt32(t *testing.T) {
	type args struct {
		num1 int32
		num2 int32
	}
	tests := []struct {
		name    string
		args    args
		wantMax int32
		wantMin int32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMax, gotMin := MaxminInt32(tt.args.num1, tt.args.num2)
			if gotMax != tt.wantMax {
				t.Errorf("MaxminInt32() gotMax = %v, want %v", gotMax, tt.wantMax)
			}
			if gotMin != tt.wantMin {
				t.Errorf("MaxminInt32() gotMin = %v, want %v", gotMin, tt.wantMin)
			}
		})
	}
}

func TestMaxminInt64(t *testing.T) {
	type args struct {
		num1 int64
		num2 int64
	}
	tests := []struct {
		name    string
		args    args
		wantMax int64
		wantMin int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMax, gotMin := MaxminInt64(tt.args.num1, tt.args.num2)
			if gotMax != tt.wantMax {
				t.Errorf("MaxminInt64() gotMax = %v, want %v", gotMax, tt.wantMax)
			}
			if gotMin != tt.wantMin {
				t.Errorf("MaxminInt64() gotMin = %v, want %v", gotMin, tt.wantMin)
			}
		})
	}
}

func TestSwap(t *testing.T) {
	type args struct {
		x int64
		y int64
	}
	tests := []struct {
		name  string
		args  args
		want  int64
		want1 int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Swap(tt.args.x, tt.args.y)
			if got != tt.want {
				t.Errorf("Swap() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Swap() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSwapString(t *testing.T) {
	type args struct {
		x string
		y string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := SwapString(tt.args.x, tt.args.y)
			if got != tt.want {
				t.Errorf("SwapString() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("SwapString() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
