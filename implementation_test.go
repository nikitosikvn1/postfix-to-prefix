package lab2

import (
    "errors"
    "fmt"
    "testing"

    "github.com/stretchr/testify/assert"
)

// Calculation correctness
func TestPostfixToPrefix1(t *testing.T) {
    res, err := PostfixToPrefix("5 6 +")
    if assert.Nil(t, err) {
        assert.Equal(t, "+ 5 6", res)
    }
}

func TestPostfixToPrefix2(t *testing.T) {
    res, err := PostfixToPrefix("2 3 ^ 4 +")
    if assert.Nil(t, err) {
        assert.Equal(t, "+ ^ 2 3 4", res)
    }
}

func TestPostfixToPrefix3(t *testing.T) {
    res, err := PostfixToPrefix("1 2 3 * + 4 +")
    if assert.Nil(t, err) {
        assert.Equal(t, "+ + 1 * 2 3 4", res)
    }
}

func TestPostfixToPrefix4(t *testing.T) {
    res, err := PostfixToPrefix("3 4 + 5 6 + *")
    if assert.Nil(t, err) {
        assert.Equal(t, "* + 3 4 + 5 6", res)
    }
}

func TestPostfixToPrefix5(t *testing.T) {
    res, err := PostfixToPrefix("2 3 + 4 * 5 - 6 / 7 8 * +")
    if assert.Nil(t, err) {
        assert.Equal(t, "+ / - * + 2 3 4 5 6 * 7 8", res)
    }
}

func TestPostfixToPrefix6(t *testing.T) {
    res, err := PostfixToPrefix("1 2 + 3 + 4 + 5 + 6 + 7 8 - 9 + 10 - *")
    if assert.Nil(t, err) {
        assert.Equal(t, "* + + + + + 1 2 3 4 5 6 - + - 7 8 9 10", res)
    }
}

// Errors
func TestPostfixToPrefixEmptyString(t *testing.T) {
    _, err := PostfixToPrefix("")
    if assert.Error(t, err) {
        assert.Equal(t, errors.New("error. Input string is empty"), err)
    }
}

func TestPostfixToPrefixInvalidString1(t *testing.T) {
    _, err := PostfixToPrefix("foo bar bazz")
    if assert.Error(t, err) {
        assert.Equal(t, errors.New("error. Invalid input string"), err)
    }
}

func TestPostfixToPrefixInvalidString2(t *testing.T) {
    _, err := PostfixToPrefix("1 2 3 * + 4 + v")
    if assert.Error(t, err) {
        assert.Equal(t, errors.New("error. Invalid input string"), err)
    }
}

func TestPostfixToPrefixInvalidString3(t *testing.T) {
    _, err := PostfixToPrefix("4 2 * 7 %")
    if assert.Error(t, err) {
        assert.Equal(t, errors.New("error. Invalid input string"), err)
    }
}

func ExamplePostfixToPrefix() {
    res, _ := PostfixToPrefix("2 2 +")
    fmt.Println(res)

    // Output:
    // + 2 2
}
