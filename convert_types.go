package mandrill

func String(v string) *string {
	return &v
}

func Int(v int) *int {
	return &v
}

func Bool(v bool) *bool {
	return &v
}

func SliceString(v []string) *[]string {
	return &v
}
