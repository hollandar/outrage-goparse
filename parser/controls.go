package parser

import (
	"github.com/hollandar/outrage-goparse/parser/matchers"
	"github.com/hollandar/outrage-goparse/parser/tokens"
)

func ControlLineFeed() matchers.Matcher {

	endOfLineTokenizer := func(value []rune) []tokens.Token {
		return []tokens.Token{tokens.NewEndOfLineToken()}
	}

	matcher := matchers.ExactMatcher{
		Match:    []rune("\r\n"),
		Tokenize: endOfLineTokenizer,
	}

	return matcher
}

func LineFeed() matchers.Matcher {

	endOfLineTokenizer := func(value []rune) []tokens.Token {
		return []tokens.Token{tokens.NewEndOfLineToken()}
	}

	matcher := matchers.ExactMatcher{
		Match:    []rune{'\n'},
		Tokenize: endOfLineTokenizer,
	}

	return matcher
}

func EndOfLine() matchers.Matcher {

	return matchers.FirstOfMatcher{Matchers: []matchers.Matcher{ControlLineFeed(), LineFeed()}}
}

func EndOfFile() matchers.Matcher {
	matcher := matchers.EndOfFileMatcher{}
	return matcher
}
