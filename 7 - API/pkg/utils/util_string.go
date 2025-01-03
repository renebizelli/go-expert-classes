package pkg_utils

import "strings"

func FormatPropertyName(input string) string {

	split := strings.Split(input, ".")

	var result []string

	for _, s := range split {
		result = append(result, strings.ToUpper(s[:1])+strings.ToLower(s[1:]))
	}

	return strings.Join(result, "")
}
