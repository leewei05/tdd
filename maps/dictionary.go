package dictionary

type Dictionary map[string]string

func (d Dictionary) Search(key string) string {
	if v, ok := d[key]; ok {
		return v
	}

	return ""
}
