package ptr

func PtrString(data string) *string {
	return &data
}

func PtrInt(data int) *int {
	return &data
}
