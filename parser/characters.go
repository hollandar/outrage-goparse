package parser

import (
	"unicode"

	"github.com/hollandar/outrage-goparse/parser/matchers"
	"github.com/hollandar/outrage-goparse/parser/tokens"
)

func stringValueTokenize(value []rune) []tokens.Token {
	svt := tokens.NewStringValueToken(value)
	tokens := []tokens.Token{svt}
	return tokens
}

func AnyChar() matchers.Matcher {
	matcher := matchers.AnyCharacterMatcher{
		Tokenize: stringValueTokenize,
	}
	return matcher
}

func Digit() matchers.Matcher {

	isDigitMatch := func(character rune) (ok bool, err string) {
		if unicode.IsDigit(character) {
			return true, ""
		} else {
			return false, "Expected a digit (0-9)."
		}
	}

	matcher := matchers.CharacterMatcher{
		Match:    isDigitMatch,
		Tokenize: stringValueTokenize,
	}
	return matcher
}

func LetterOrDigit() matchers.Matcher {

	isLetterOrDigitMatch := func(character rune) (ok bool, err string) {
		if unicode.IsLetter(character) || unicode.IsDigit(character) {
			return true, ""
		} else {
			return false, "Expected a digit (a-zA-Z0-9)."
		}
	}

	matcher := matchers.CharacterMatcher{
		Match:    isLetterOrDigitMatch,
		Tokenize: stringValueTokenize,
	}
	return matcher
}

func Letter() matchers.Matcher {

	isLetterMatch := func(character rune) (ok bool, err string) {
		if unicode.IsLetter(character) {
			return true, ""
		} else {
			return false, "Expected a letter (a-zA-Z)."
		}
	}

	matcher := matchers.CharacterMatcher{
		Match:    isLetterMatch,
		Tokenize: stringValueTokenize,
	}
	return matcher
}

func UppercaseLetter() matchers.Matcher {

	isUppercaseLetterMatch := func(character rune) (ok bool, err string) {
		if unicode.IsLetter(character) && unicode.IsUpper(character) {
			return true, ""
		} else {
			return false, "Expected a letter (A-Z)."
		}
	}

	matcher := matchers.CharacterMatcher{
		Match:    isUppercaseLetterMatch,
		Tokenize: stringValueTokenize,
	}
	return matcher
}

func LowercaseLetter() matchers.Matcher {

	isLowercaseLetterMatch := func(character rune) (ok bool, err string) {
		if unicode.IsLetter(character) && unicode.IsUpper(character) {
			return true, ""
		} else {
			return false, "Expected a letter (a-z)."
		}
	}

	matcher := matchers.CharacterMatcher{
		Match:    isLowercaseLetterMatch,
		Tokenize: stringValueTokenize,
	}
	return matcher
}

func Period() matchers.Matcher                { return Char('.') }
func Comma() matchers.Matcher                 { return Char(',') }
func Semicolon() matchers.Matcher             { return Char(';') }
func LeftBracket() matchers.Matcher           { return Char('(') }
func RightBracket() matchers.Matcher          { return Char(')') }
func LeftSquareBracket() matchers.Matcher     { return Char('[') }
func RightSquareBracket() matchers.Matcher    { return Char(']') }
func LeftBrace() matchers.Matcher             { return Char('{') }
func RightBrace() matchers.Matcher            { return Char('}') }
func ForwardSlash() matchers.Matcher          { return Char('/') }
func Divide() matchers.Matcher                { return ForwardSlash() }
func BackSlash() matchers.Matcher             { return Char('\\') }
func Underscore() matchers.Matcher            { return Char('_') }
func Plus() matchers.Matcher                  { return Char('+') }
func Add() matchers.Matcher                   { return Plus() }
func Minus() matchers.Matcher                 { return Char('-') }
func Subtract() matchers.Matcher              { return Minus() }
func Equal() matchers.Matcher                 { return Char('=') }
func Equality() matchers.Matcher              { return String("==") }
func NotEquality() matchers.Matcher           { return String("!=") }
func LessThan() matchers.Matcher              { return Char('<') }
func LessThanOrEquality() matchers.Matcher    { return String("<=") }
func GreaterThan() matchers.Matcher           { return Char('>') }
func GreaterThanOrEquality() matchers.Matcher { return String(">=") }
func Ampersand() matchers.Matcher             { return Char('&') }
func Multiply() matchers.Matcher              { return Char('*') }
func Asterisk() matchers.Matcher              { return Multiply() }
func Space() matchers.Matcher                 { return Char(' ') }
func Tab() matchers.Matcher                   { return Char('\t') }
func Whitespace() matchers.Matcher            { return Or(Space(), Tab()) }
func Whitespaces() matchers.Matcher           { return Some(Whitespace()) }
