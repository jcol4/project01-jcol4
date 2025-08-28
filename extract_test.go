package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestExtract(t *testing.T) {
	tests := []struct {
		words, hrefs []string
		bodyText     []byte
	}{
		{[]string{"hello", "world"}, []string{"https://example.com"}, []byte("<p>hello world</p>")},
	}
	for _, test := range tests {
		wordRes, hrefRes := Extract(test.bodyText)
		if !cmp.Equal(wordRes, test.words) || !cmp.Equal(hrefRes, test.hrefs) {
			t.Errorf("Expected Words: %s HREF %s, got Words %s HREF %s", test.words, test.hrefs, wordRes, hrefRes)
		}
	}
}
