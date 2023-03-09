package lab2

import (
    "bytes"
    "errors"
    "io"
    "strings"
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestComputeHandler_Compute(t *testing.T) {
    type testCase struct {
        input           string
        expectedOutput  string
        expectedError   error
    }

    testCases := []testCase{
        {
            input:          "2 3 +",
            expectedOutput: "+ 2 3",
            expectedError:  nil,
        },
        {
            input:          "4 5 * 3 -",
            expectedOutput: "- * 4 5 3",
            expectedError:  nil,
        },
		{
            input:          "10 5 - 2 / 3 4 * +",
            expectedOutput: "+ / - 10 5 2 * 3 4",
            expectedError:  nil,
        },
		{
            input:          "4 5 + 6 *",
            expectedOutput: "* + 4 5 6",
            expectedError:  nil,
        },
        {
            input:          "",
            expectedOutput: "",
            expectedError:  io.EOF,
        },
        {
            input:          "invalid input",
            expectedOutput: "",
            expectedError:  errors.New("error. Invalid input string"),
        },
    }

    for _, tc := range testCases {
        reader := strings.NewReader(tc.input)
        writer := &bytes.Buffer{}
        ch := NewComputeHandler(reader, writer)

        err := ch.Compute()

        assert.Equal(t, tc.expectedError, err)
        assert.Equal(t, tc.expectedOutput, writer.String())
    }
}