package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestExamples(t *testing.T) {
	tests := []struct {
		env    map[string]string
		input  string
		output string
	}{
		{map[string]string{
			"ANSWER": "42",
			"ITEM":   "Towel",
		}, `
{{ range $key, $value := . }}
  KEY:{{ $key }} VALUE:{{ $value }}
{{ end }}`, `

  KEY:ANSWER VALUE:42

  KEY:ITEM VALUE:Towel
`},
	}
	var buf bytes.Buffer
	for i, d := range tests {
		//t.Log(i, d)
		buf.Reset()
		err := tmpl(strings.NewReader(d.input), &buf, d.env)
		if err != nil {
			t.Errorf("%2d failed: %v", i, err)
			continue
		}
		if g, w := buf.String(), d.output; g != w {
			t.Errorf("%2d\t got %q,\n\t\twant %q", i, g, w)
		}
	}
}
