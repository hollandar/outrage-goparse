package calctokens

import (
	parseTokens "github.com/hollandar/outrage-goparse/parser/tokens"
)

type DecimalToken struct {
	parseTokens.Token
	Value float64
}
