package dictionary

func Search(dic map[string]string, key string) string {
	if v, ok := dic[key]; ok {
		return v
	}

	return ""
}
