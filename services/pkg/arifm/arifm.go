package arifm

import shuntingYard "github.com/mgenware/go-shunting-yard"

func ArifmeticServer(expression string) (int, error) {

	infinxToken, err := shuntingYard.Scan(expression)
	if err != nil {
		return -1, err
	}

	postfixToken, err := shuntingYard.Parse(infinxToken)
	if err != nil {
		return -1, err
	}

	result, err := shuntingYard.Evaluate(postfixToken)
	if err != nil {
		return -1, err
	}
	return result, nil
}
