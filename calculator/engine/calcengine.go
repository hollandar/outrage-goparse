package engine

import (
	"errors"
	"math"

	"github.com/hollandar/outrage-goparse/calculator/calctokens"
	"github.com/hollandar/outrage-goparse/parser/tokens"
)

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

type OpFunc func(float64, float64) float64

func processOperation[T interface{}](list []tokens.Token, op OpFunc) ([]tokens.Token, error) {

	finished := false

	for !finished {
		found := false
		for i := 0; i < len(list); i++ {
			_, ok := list[i].(T)
			if ok {
				if i <= 0 || i >= len(list)-1 {
					return nil, errors.New("Operation was out of place.")
				}

				left, leftOk := list[i-1].(calctokens.DecimalToken)
				right, rightOk := list[i+1].(calctokens.DecimalToken)

				if !leftOk || !rightOk {
					return nil, errors.New("Decimals out of place around operation.")
				}

				leftValue := left.Value
				rightValue := right.Value

				result := op(leftValue, rightValue)

				newList := append(list[:i-1], calctokens.DecimalToken{Value: result})
				newList = append(newList, list[i+2:]...)

				list = newList
				found = true
				break
			}
		}
		if !found {
			finished = true
		}
	}

	return list, nil
}
