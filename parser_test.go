package argsparsego

import (
	"strings"
	"testing"
)

func argsArrayString(l []string) string {
	var builder strings.Builder
	for i, v := range l {
		if i != 0 {
			builder.WriteString(", ")
		}
		builder.WriteString("\"")
		builder.WriteString(v)
		builder.WriteString("\"")
	}
	return builder.String()
}

type testCase struct {
	input    string
	expected []string
}

func runTests(t *testing.T, tests []testCase) {
	for _, test := range tests {
		l, err := Parse(test.input)
		if err != nil {
			t.Errorf("Parse(%s) error %v", test.input, err)
		}
		if len(l) != len(test.expected) {
			t.Errorf("Parse(%s) = [ %s ], want [ %s ]", test.input, argsArrayString(l), argsArrayString(test.expected))
		}
		for i, v := range l {
			if v != test.expected[i] {
				t.Errorf("Parse(%s) = %v, want %v", test.input, argsArrayString(l), argsArrayString(test.expected))
			}
		}
	}
}

func TestParse(t *testing.T) {
	tests := []testCase{
		{"hoge fuga", []string{"hoge", "fuga"}},
	}
	runTests(t, tests)
}

func TestParseQuote(t *testing.T) {
	tests := []testCase{
		{"hoge 'fuga hoge'", []string{"hoge", "fuga hoge"}},
	}
	runTests(t, tests)
}

func TestParseDoubleQuote(t *testing.T) {
	tests := []testCase{
		{"hoge \"fuga hoge\"", []string{"hoge", "fuga hoge"}},
	}
	runTests(t, tests)
}

func TestParseEscape(t *testing.T) {
	tests := []testCase{
		{"hoge\\ fuga", []string{"hoge fuga"}},
	}
	runTests(t, tests)
}

func TestParseNestedDoubleQuoteAndEscape(t *testing.T) {
	tests := []testCase{
		{"hoge \"fuga\\ hoge\"", []string{"hoge", "fuga hoge"}},
		{"hoge \"fuga\\\" hoge\"", []string{"hoge", "fuga\" hoge"}},
	}
	runTests(t, tests)
}

func TestParseNestedQuoteAndEscape(t *testing.T) {
	tests := []testCase{
		{"hoge 'fuga\\ hoge'", []string{"hoge", "fuga hoge"}},
		{"hoge 'fuga\\' hoge'", []string{"hoge", "fuga' hoge"}},
	}
	runTests(t, tests)
}

func TestParseNestedQuoteAndDoubleQuote(t *testing.T) {
	tests := []testCase{
		{"hoge 'fuga \"hoge\"'", []string{"hoge", "fuga \"hoge\""}},
		{"hoge \"fuga 'hoge'\"", []string{"hoge", "fuga 'hoge'"}},
	}
	runTests(t, tests)
}
