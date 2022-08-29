package parse

import (
	"errors"
	"strconv"

	"github.com/hollandar/outrage-goparse/calculator/calctokens"
	"github.com/hollandar/outrage-goparse/parser"
	"github.com/hollandar/outrage-goparse/parser/matchers"
	"github.com/hollandar/outrage-goparse/parser/tokens"
)

func Decimal() matchers.Matcher {

	toDecimalToken := func(value string) ([]tokens.Token, error) {
		val, error := strconv.ParseFloat(value, 64)
		if error != nil {
			return []tokens.Token{}, error
		}

		token := calctokens.DecimalToken{
			Value: val,
		}

		return []tokens.Token{token}, nil
	}

	return parser.Convert(
		parser.FirstOf(
			parser.Sequence(parser.Some(parser.Digit()), parser.Period(), parser.Some(parser.Digit())),
			parser.Then(parser.Period(), parser.Some(parser.Digit())),
			parser.Some(parser.Digit()),
		),
		toDecimalToken,
	)
}

func Sqrt() matchers.Matcher {
	return parser.Wrap(
		parser.Sequence(
			parser.Ignore(parser.String("sqrt")),
			parser.Surrounded(parser.Ref(Expression), parser.Ignore(parser.LeftBracket()), parser.Ignore(parser.RightBracket())),
		),
		&calctokens.FunctionToken{
			Function: calctokens.Function_Sqrt,
		},
	)
}

func Functions() matchers.Matcher {
	return parser.FirstOf(
		Sqrt(),
	)
}

func Raise() matchers.Matcher {
	return parser.Produce(
		parser.Char('^'),
		calctokens.RaiseToken{},
	)
}

func Add() matchers.Matcher {
	return parser.Produce(
		parser.Add(),
		calctokens.AddToken{},
	)
}

func Subtract() matchers.Matcher {
	return parser.Produce(
		parser.Subtract(),
		calctokens.SubtractToken{},
	)
}

func Multiply() matchers.Matcher {
	return parser.Produce(
		parser.Multiply(),
		calctokens.MultiplyToken{},
	)
}

func Divide() matchers.Matcher {
	return parser.Produce(
		parser.Divide(),
		calctokens.DivideToken{},
	)
}

func Expression() matchers.Matcher {
	return parser.Many(
		parser.FirstOf(
			Functions(),
			Raise(),
			Add(),
			Subtract(),
			Multiply(),
			Divide(),
			Decimal(),
			BracketedExpression(),
			parser.Ignore(parser.Whitespaces()),
		),
	)
}

func BracketedExpression() matchers.Matcher {
	return parser.Wrap(
		parser.Surrounded(
			parser.Ref(Expression),
			parser.Ignore(parser.LeftBracket()),
			parser.Ignore(parser.RightBracket()),
		),
		&calctokens.BracketsToken{},
	)
}

func Calculation() matchers.Matcher {
	return parser.Then(
		Expression(),
		parser.Ignore(parser.EndOfFile()),
	)
}

func ParseExpression(expression string) ([]tokens.Token, error) {
	match := parser.Parse(expression, Calculation())
	if match.Success {
		return match.Tokens, nil
	} else {
		return nil, errors.New(match.Error)
	}
}
