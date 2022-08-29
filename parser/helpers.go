package parser

import (
	"fmt"
	"math"
	"strings"

	"github.com/hollandar/outrage-goparse/parser/matchers"
	"github.com/hollandar/outrage-goparse/parser/tokens"
)

func Ignore(ignore matchers.Matcher) matchers.Matcher {
	matcher := matchers.IgnoreMatcher{
		Ignore: ignore,
	}

	return matcher
}

func Cast(innerMatcher matchers.Matcher, castTo tokens.Token) matchers.Matcher {

	castFunc := func(from tokens.Token) tokens.Token {
		return castTo
	}

	matcher := matchers.CastMatcher{
		InnerMatcher: innerMatcher,
		Cast:         castFunc,
	}

	return matcher
}

func Wrap(innerMatcher matchers.Matcher, wrapper tokens.Wrapper) matchers.Matcher {

	wrapFunc := func(match matchers.Match) ([]tokens.Token, error) {
		wrapper.SetTokens(match.Tokens)

		return []tokens.Token{wrapper}, nil
	}

	matcher := matchers.WrapMatcher{
		Inner: innerMatcher,
		Wrap:  wrapFunc,
	}

	return matcher
}

func Identifier(innerMatcher matchers.Matcher) matchers.Matcher {

	wrapFunc := func(match matchers.Match) ([]tokens.Token, error) {

		result := ""
		for _, token := range match.Tokens {
			s, ok := token.(tokens.StringValueToken)
			if ok {
				result += string(s.Value)
			}
		}

		return []tokens.Token{tokens.NewIdentifierToken(result)}, nil

	}
	matcher := matchers.WrapMatcher{
		Inner: innerMatcher,
		Wrap:  wrapFunc,
	}

	return matcher
}

func Char(character rune) matchers.Matcher {
	isChar := func(c rune) (ok bool, err string) {
		if c == character {
			return true, ""
		} else {
			return false, fmt.Sprintf("Expected character %c", character)
		}
	}

	matcher := matchers.CharacterMatcher{
		Match:    isChar,
		Tokenize: stringValueTokenize,
	}
	return matcher
}

func String(value string) matchers.Matcher {
	matcher := matchers.ExactMatcher{
		Match:    []rune(value),
		Tokenize: stringValueTokenize,
	}

	return matcher
}

func Text(innerMatcher matchers.Matcher) matchers.Matcher {

	wrapFunc := func(match matchers.Match) ([]tokens.Token, error) {

		result := ""
		for _, token := range match.Tokens {
			s, ok := token.(tokens.StringValueToken)
			if ok {
				result += string(s.Value)
			}
		}

		return []tokens.Token{tokens.NewTextToken(result)}, nil

	}
	matcher := matchers.WrapMatcher{
		Inner: innerMatcher,
		Wrap:  wrapFunc,
	}

	return matcher
}

func TextTrim(innerMatcher matchers.Matcher, characters ...rune) matchers.Matcher {

	wrapFunc := func(match matchers.Match) ([]tokens.Token, error) {

		result := ""
		for _, token := range match.Tokens {
			s, ok := token.(tokens.StringValueToken)
			if ok {
				result += string(s.Value)
			}
		}

		result = strings.Trim(result, string(characters))

		return []tokens.Token{tokens.NewTextToken(result)}, nil

	}
	matcher := matchers.WrapMatcher{
		Inner: innerMatcher,
		Wrap:  wrapFunc,
	}

	return matcher
}

func TextTrimWhitespace(innerMatcher matchers.Matcher) matchers.Matcher {
	return TextTrim(innerMatcher, ' ')
}

type ConvertFunc func(value string) ([]tokens.Token, error)

func Convert(innerMatcher matchers.Matcher, convertFunc ConvertFunc) matchers.Matcher {

	wrapFunc := func(match matchers.Match) ([]tokens.Token, error) {

		result := ""
		for _, token := range match.Tokens {
			s, ok := token.(tokens.StringValueToken)
			if ok {
				result += string(s.Value)
			}
		}

		return convertFunc(result)
	}

	matcher := matchers.WrapMatcher{
		Inner: innerMatcher,
		Wrap:  wrapFunc,
	}

	return matcher
}

func Produce(innerMatcher matchers.Matcher, produce tokens.Token) matchers.Matcher {

	castFunc := func(from tokens.Token) tokens.Token {
		return produce
	}

	matcher := matchers.CastMatcher{
		InnerMatcher: innerMatcher,
		Cast:         castFunc,
	}

	return matcher
}

func Then(first matchers.Matcher, then matchers.Matcher) matchers.Matcher {
	matcher := matchers.ThenMatcher{
		InitialMatcher: first,
		ThenMatcher:    then,
	}

	return matcher
}

func Sequence(first matchers.Matcher, second matchers.Matcher, other ...matchers.Matcher) matchers.Matcher {
	matchers := []matchers.Matcher{first, second}
	matchers = append(matchers, other...)

	lastTwo := matchers[len(matchers)-2:]

	matcher := Then(lastTwo[0], lastTwo[1])
	if len(matchers) > 2 {
		others := matchers[:len(matchers)-2]

		for i := len(others) - 1; i >= 0; i-- {
			matcher = Then(others[i], matcher)
		}
	}

	return matcher
}

func Many(inner matchers.Matcher) matchers.Matcher {
	matcher := matchers.ManyMatcher{
		Many:           inner,
		MinimumMatches: 0,
		MaximumMatches: math.MaxInt,
	}

	return matcher
}

func Once(inner matchers.Matcher) matchers.Matcher {
	matcher := matchers.ManyMatcher{
		Many:           inner,
		MinimumMatches: 1,
		MaximumMatches: 1,
	}

	return matcher
}

func Optional(inner matchers.Matcher) matchers.Matcher {
	matcher := matchers.ManyMatcher{
		Many:           inner,
		MinimumMatches: 0,
		MaximumMatches: 1,
	}

	return matcher
}

func Some(many matchers.Matcher) matchers.Matcher {
	matcher := matchers.ManyMatcher{
		Many:           many,
		MinimumMatches: 1,
		MaximumMatches: math.MaxInt,
	}

	return matcher
}

func SomeThen(many matchers.Matcher, then matchers.Matcher) matchers.Matcher {
	return CountThen(many, then, 1)
}

func ManyThen(many matchers.Matcher, then matchers.Matcher) matchers.Matcher {
	return CountThen(many, then, 0)
}

func CountThen(many matchers.Matcher, then matchers.Matcher, minimumMatches int) matchers.Matcher {
	matcher := matchers.ManyThenMatcher{
		Many:           many,
		Then:           then,
		MinimumMatches: minimumMatches,
	}

	return matcher
}

func FirstOf(inner ...matchers.Matcher) matchers.Matcher {
	matcher := matchers.FirstOfMatcher{
		Matchers: inner,
	}

	return matcher
}

func Block(inner matchers.Matcher) matchers.Matcher {
	matcher := matchers.BlockMatcher{
		InnerMatcher: inner,
	}

	return matcher
}

func Preview(inner matchers.Matcher) matchers.Matcher {
	matcher := matchers.PreviewMatcher{
		Preview: inner,
	}

	return matcher
}

func Or(matcherList ...matchers.Matcher) matchers.Matcher {
	matcher := matchers.FirstOfMatcher{
		Matchers: matcherList,
	}

	return matcher
}

func DelimitedBy(input matchers.Matcher, delimiter matchers.Matcher, minOccurrence int, maxOccurrence int) matchers.Matcher {
	matcher := matchers.DelimitedByMatcher{
		Term:          input,
		Delimiter:     delimiter,
		MinOccurrence: minOccurrence,
		MaxOccurrence: maxOccurrence,
	}

	return matcher
}

func Surrounded(innerMatcher, leftMatcher, rightMatcher matchers.Matcher) matchers.Matcher {
	return Then(leftMatcher, Until(innerMatcher, rightMatcher))
}

func Until(input matchers.Matcher, until matchers.Matcher) matchers.Matcher {
	matcher := matchers.UntilMatcher{
		Matcher: input,
		Until:   until,
	}

	return matcher
}

func When(innerMatcher matchers.Matcher, clause matchers.WhenFunc, msg string) matchers.Matcher {
	matcher := matchers.WhenMatcher{
		Initial: innerMatcher,
		Clause:  clause,
		Msg:     msg,
	}

	return matcher
}

func Except(actual matchers.Matcher, except matchers.Matcher) matchers.Matcher {
	matcher := matchers.ExceptMatcher{
		Except: except,
		Actual: actual,
	}

	return matcher
}

func Ref(ref matchers.RefFunc) matchers.Matcher {
	matcher := matchers.RefMatcher{Ref: ref}
	return matcher
}
