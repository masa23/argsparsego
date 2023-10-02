package argsparsego

import (
	"errors"
	"strings"
)

func Parse(argsString string) ([]string, error) {
	var args []string
	var currentArg strings.Builder
	var inSingleQuote bool
	var inDoubleQuote bool
	var escapeNext bool

	for _, r := range argsString {
		switch {
		case escapeNext:
			currentArg.WriteRune(r)
			escapeNext = false
		case r == '\\':
			escapeNext = true
		case r == '"' && !inSingleQuote:
			if inDoubleQuote && currentArg.Len() > 0 {
				args = append(args, currentArg.String())
				currentArg.Reset()
			}
			inDoubleQuote = !inDoubleQuote
		case r == '\'' && !inDoubleQuote:
			if inSingleQuote && currentArg.Len() > 0 {
				args = append(args, currentArg.String())
				currentArg.Reset()
			}
			inSingleQuote = !inSingleQuote
		case r == ' ' && !inSingleQuote && !inDoubleQuote:
			if currentArg.Len() > 0 {
				args = append(args, currentArg.String())
				currentArg.Reset()
			}
		default:
			currentArg.WriteRune(r)
		}
	}

	if inSingleQuote || inDoubleQuote {
		return nil, errors.New("unmatched quote in input string")
	}

	if currentArg.Len() > 0 {
		args = append(args, currentArg.String())
	}

	return args, nil
}
