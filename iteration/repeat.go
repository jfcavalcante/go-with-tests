package iteration

import "strings"

func Repeat(s string, count int) string {
	var repeated string
	for i := 0; i < count; i++ {
		repeated = strings.Clone(s) + repeated
	}
	return repeated
}
