package utils

import (
	"fmt"
	"strings"
)

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

// commonInitialisms from https://github.com/golang/lint/blob/master/lint.go#L770
var commonInitialisms = []string{"ACL", "API", "ASCII", "CPU", "CSS", "DNS", "EOF", "GUID", "HTML", "HTTP", "HTTPS", "ID", "IP", "JSON", "LHS", "QPS", "RAM", "RHS", "RPC", "SLA", "SMTP", "SQL", "SSH", "TCP", "TLS", "TTL", "UDP", "UI", "UID", "UUID", "URI", "URL", "UTF8", "VM", "XML", "XMPP", "XSRF", "XSS"}

//CommonInitialisms 替换专有名词
func CommonInitialisms(s string) string {
	var commonInitialismsReplacer []string
	for i := range commonInitialisms {
		initialism := commonInitialisms[i]
		l := strings.ToLower(initialism)
		commonInitialismsReplacer = append(commonInitialismsReplacer, strings.ToUpper(l[:1])+l[1:], initialism)
	}
	return strings.NewReplacer(commonInitialismsReplacer...).Replace(s)
}
