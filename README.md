# Outrage GoParse

GoParse is a port of the equivalent C# parser to GO, with improvements in the api so it is more usable with golang.

The parser uses a tree of parser functions that are evaluated from front to back, in order to parse a text based language into a collection of typed structs which you can introspet as you interpret the tokens returned by the parser.

## Getting started

Include a dependency on the parsing module:

``` go
  import "github.com/hollandar/outrage-goparse/parser"
```

### Parsing

Define a parser, below is an example of the decimal parser from the calculator example in this repo at /calculator/parser.go.

``` go
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
```

This parser matches first full decimals (00.00) then leading period values (.00) then whole integers (00), the creates a token out of the matched value, using an inline func to convert the parsed string into a decimal.
Failures during the conversion to a token can be returned as errors, which end up in the parsers core error list.

### Circular parsing

If your parsing functions need a circular dependency, a parser refers to another parsing function that refers back to the original, you can use parse.Ref to construct the parser.  See this example from the calculator for bracketed expressions.

```go
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
```

Because Expression refers to BracketedExpression, parser.Ref is used to decouple the circular dependency.  If you receive a stack overflow during parsing at runtime, it is likely you have a circular dependency in your parser.

### Tokenizing

`parser.Convert` is responsible for consuming the parsed string value and converting it into a token of the correct type.  A token is your own struct which references the Token interface as follows:

``` go
package calctokens

import (
	parseTokens "github.com/hollandar/outrage-goparse/parser/tokens"
)

type DecimalToken struct {
	parseTokens.Token
	Value float64
}
```
The inclusion of the Token interface is the only imposition, so your token can be tracked in the eventual heirarchy of tokens.

### Initiation

The chain of functions that perform the parsing is heirarchical and must begin at a root parser which you pass to Parse along with the expression (string) being parsed.

``` go
func ParseExpression(expression string) ([]tokens.Token, error) {
	match := parser.Parse(expression, Calculation())
	if match.Success {
		return match.Tokens, nil
	} else {
		return nil, errors.New(match.Error)
	}
}
```

This process will return a Match object the contains a list of Tokens that represents the root objects in your token tree (or parsing errors), which you can then interpret using an interpreter or a code generator.

### Interpreter

The calculator interpreter is a fairy simple example of an interpreter.  It recursively calculates the expressions passed to functions, then bracketed expressions, then other operators according to the usual rules.
The result of the parsing is a tree of tokens for bracketed expressions or token lists for normal expressions.
"1 + 5" will result in three tokens DecimalToken(1) AddToken DecimalToken(5), so the approach taken is to interpret this triplet and replace it with DecimalToken(6).
"1 + (5)" on the other hand is recursive, tokens for this expressions are DecimalToken(1) AddToken BracketedToken[DecimalToken(5)].  The approach here is to handle the bracketed token and replace it with DecimalToken(5) by interpreting its internals, then proceeding as above.

```go
func Calculate(list []tokens.Token) (float64, error) {
	for i := 0; i < len(list); i++ {
		currentToken := list[i]

		bracketsToken, ok := currentToken.(*calctokens.BracketsToken)
		if ok {
			bracketsValue, err := Calculate(bracketsToken.Value)
			if err != nil {
				return 0, err
			}
			list[i] = calctokens.DecimalToken{Value: bracketsValue}
		}

		functionToken, ok := currentToken.(*calctokens.FunctionToken)
		if ok {
			functionValue, err := Calculate(functionToken.Parameters)
			if err != nil {
				return 0, err
			}
			switch functionToken.Function {
			case calctokens.Function_Sqrt:
				functionValue = math.Sqrt(functionValue)
			}

			list[i] = calctokens.DecimalToken{Value: functionValue}
		}
	}

	raiseOp := func(a, b float64) float64 { return math.Pow(a, b) }
	list, err := processOperation[calctokens.RaiseToken](list, raiseOp)
	if err != nil {
		return 0, err
	}
	divideOp := func(a, b float64) float64 { return a / b }
	list, err = processOperation[calctokens.DivideToken](list, divideOp)
	if err != nil {
		return 0, err
	}

	multiplyOp := func(a, b float64) float64 { return a * b }
	list, err = processOperation[calctokens.MultiplyToken](list, multiplyOp)
	if err != nil {
		return 0, err
	}

	addOp := func(a, b float64) float64 { return a + b }
	list, err = processOperation[calctokens.AddToken](list, addOp)
	if err != nil {
		return 0, err
	}

	subtractOp := func(a, b float64) float64 { return a - b }
	list, err = processOperation[calctokens.SubtractToken](list, subtractOp)
	if err != nil {
		return 0, err
	}

	if len(list) != 1 {
		return 0, errors.New("Operations incomplete")
	}

	decimalToken, ok := list[0].(calctokens.DecimalToken)
	if !ok {
		return 0, errors.New("Did not end up with a decimal.")
	}

	return decimalToken.Value, nil
}
```

This library works correctly, but its tests have not been converted yet.  Happy to take your feedback!
