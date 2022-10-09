package utils

func Underline2UpperCamelCase(s string) string {
	var result string
	for i, v := range s {
		if v == '_' {
			continue
		}
		if i == 0 || (i > 0 && s[i-1] == '_') {
			result += string(v - 32)
		} else {
			result += string(v)
		}
	}
	return result
}
