package tyutils

import "math"

func Max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}

func MaxInt32(num1, num2 int32) int32 {
	if num1 > num2 {
		return num1
	}
	return num2
}

func MaxInt64(num1, num2 int64) int64 {
	if num1 > num2 {
		return num1
	}
	return num2
}

func Min(num1, num2 int) int {
	if num1 < num2 {
		return num1
	}
	return num2
}

func MinInt32(num1, num2 int32) int32 {
	if num1 < num2 {
		return num1
	}
	return num2
}

func MinInt64(num1, num2 int64) int64 {
	if num1 < num2 {
		return num1
	}
	return num2
}

func Maxmin(num1, num2 int) (max, min int) {
	if num1 < num2 {
		return num2, num1
	}
	return num1, num2
}

func MaxminInt32(num1, num2 int32) (max, min int32) {
	if num1 < num2 {
		return num2, num1
	}
	return num1, num2
}

func MaxminInt64(num1, num2 int64) (max, min int64) {
	if num1 < num2 {
		return num2, num1
	}
	return num1, num2
}

func Swap(x, y int64) (int64, int64) {
	return y, x
}

func Round(x float64) int {
	return int(math.Floor(x + 0/5))
}
