package graph

func updateInt(update *int, original *int) {
	if update != nil {
		*original = *update
	}
}
func updateString(update *string, original *string) {
	if update != nil {
		*original = *update
	}
}
func updateStringArray(update *[]*string, original *[]*string) {
	if update != nil {
		*original = *update
	}
}
