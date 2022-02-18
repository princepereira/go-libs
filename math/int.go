package math

// MaxInt returns the maximum int value from list
func MaxInt(a int, b ...int) int {
	for _, v := range b {
		if v > a {
			a = v
		}
	}
	return a
}

// MinInt returns the minimum int value from list provided
func MinInt(a int, b ...int) int {
	for _, v := range b {
		if v < a {
			a = v
		}
	}
	return a
}

// MaxInt32 returns the maximum int32 value from list
func MaxInt32(a int32, b ...int32) int32 {
	for _, v := range b {
		if v > a {
			a = v
		}
	}
	return a
}

// MinInt32 returns the minimum int32 value from list provided
func MinInt32(a int32, b ...int32) int32 {
	for _, v := range b {
		if v < a {
			a = v
		}
	}
	return a
}

// MaxInt64 returns the maximum int64 value from list
func MaxInt64(a int64, b ...int64) int64 {
	for _, v := range b {
		if v > a {
			a = v
		}
	}
	return a
}

// MinInt64 returns the minimum int64 value from list provided
func MinInt64(a int64, b ...int64) int64 {
	for _, v := range b {
		if v < a {
			a = v
		}
	}
	return a
}
