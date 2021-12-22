package conv

func PointerString(s string) *string {
	return &s
}

func PointerInt64(i int64) *int64 {
	return &i
}
