package utils

import "fmt"

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
func PrintGreen(s string) {
	fmt.Printf("\033[32m%s\033[0m\n", s)
}

func PrintYellow(s string) {
	fmt.Printf("\033[33m%s\033[0m\n", s)
}
func PrintRed(s string) {
	fmt.Printf("\033[31m%s\033[0m\n", s)
}
