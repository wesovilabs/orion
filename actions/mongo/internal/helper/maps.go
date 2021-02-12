package helper

func MapStructToArray(dict map[string]struct{}) []string {
	out := make([]string, len(dict))
	index := 0
	for name := range dict {
		out[index] = name
		index++
	}
	return out
}
