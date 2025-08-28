package main

import "strings"

func Clean(host *string, hrefs *[]string) []string {
	cleanUrls := []string{*host}
	for _, str := range *hrefs {
		if str != "/" {
			if !strings.Contains(str, "https://") {
				var s strings.Builder
				s.WriteString(strings.TrimSuffix(*host, "/") + str)
				cleanUrls = append(cleanUrls, s.String())
			} else {
				cleanUrls = append(cleanUrls, str)
			}
		}
	}
	return cleanUrls
}
