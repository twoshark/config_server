package graph

func updateInt(update *int, original *int) *int {
	if update != nil {
		return original
	}
	return update
}
func updateString(update *string, original *string) *string {
	if update != nil {
		return original
	}
	return update
}
func updateStringArray(update []*string, original []*string) []*string {
	if update != nil {
		return original
	}
	return update
}
