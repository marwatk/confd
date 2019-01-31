package template_test

import (
	"testing"

	"github.com/kelseyhightower/confd/resource/template"
)

func TestEscapeOsgi(t *testing.T) {
	// quotes, double quotes, backslash, the equals sign and spaces need to be escaped
	tests := [][]string{
		[]string{"", ""},
		[]string{"a", "a"},
		[]string{"'", "\\'"},
		[]string{"\"", "\\\""},
		[]string{"\\", "\\\\"},
		[]string{"=", "\\="},
		[]string{" ", "\\ "},
		[]string{
			"a long 'sentence' using \"some\" of the = characters",
			"a\\ long\\ \\'sentence\\'\\ using\\ \\\"some\\\"\\ of\\ the\\ \\=\\ characters",
		},
	}
	for _, test := range tests {
		if actual := template.EscapeOsgi(test[0]); actual != test[1] {
			t.Errorf("EscapeOsgi failed: [%s] != [%s]", actual, test[1])
		}
	}
}
