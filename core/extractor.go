package core

import "regexp"

func Extractor(html string) string {
	re := regexp.MustCompile(`var\s+ytInitialData\s*=\s*({.+?});`)
	matches := re.FindAllStringSubmatch(html, -1)
	if len(matches) > 0 {
		return matches[0][1]
	}
	return ""
}
