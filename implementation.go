package lab2

import (
    "errors"
    "fmt"
    "strings"
)

// Function to convert postfix to prefix form
func PostfixToPrefix(postfix string) (string, error) {
    if postfix == "" {
        return "", errors.New("error. Input string is empty")
    }

    stack := make([]string, 0)
    tokens := strings.Split(postfix, " ")

    for _, token := range tokens {
        if isOperand(token) {
            stack = append(stack, token)
        } else {
            if len(stack) < 2 {
                return "", errors.New("error. Invalid input string")
            }

            operand1 := stack[len(stack)-1]
            operand2 := stack[len(stack)-2]
            stack = stack[:len(stack)-2]

            newToken := fmt.Sprintf("%s %s %s", token, operand2, operand1)
            stack = append(stack, newToken)
        }
    }

    if len(stack) != 1 {
        return "", errors.New("error. Invalid input string")
    }

    return stack[0], nil
}


// The function that checks the operand
func isOperand(token string) bool {
    switch token {
    case "+", "-", "*", "/", "^":
        return false
    default:
        return true
    }
}