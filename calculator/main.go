package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/hollandar/outrage-goparse/calculator/engine"
	"github.com/hollandar/outrage-goparse/calculator/parse"
)

func main() {
	for {

		fmt.Print("calc >")
		var input string
		//input = "10 + sqrt(16)"

		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			break
		}

		input = strings.TrimSuffix(input, "\r\n")

		if input == "exit" {
			break
		}

		tokens, err := parse.ParseExpression(input)

		if err != nil {
			fmt.Println(err.Error())
		} else {
			result, err := engine.Calculate(tokens)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println(result)
			}
		}
	}

	fmt.Println("Done.")
}
