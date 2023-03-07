package lab2

import (
	"bufio"
	"io"
)

// ComputeHandler should be constructed with input io.Reader and output io.Writer.
// Its Compute() method should read the expression from input and write the computed result to the output.

type ComputeHandler struct {
    reader io.Reader
    writer io.Writer
}

func NewComputeHandler(reader io.Reader, writer io.Writer) *ComputeHandler {
    return &ComputeHandler{
        reader: reader,
        writer: writer,
    }
}

func (ch *ComputeHandler) Compute() error {
	// Read the expression from the input reader
    scanner := bufio.NewScanner(bufio.NewReader(ch.reader))
    if !scanner.Scan() {
        if scanner.Err() != nil {
            return scanner.Err()
        }
        return io.EOF
    }
    postfixExpr := scanner.Text()

	// Convert postfix to prefix
    prefixExpr, err := PostfixToPrefix(postfixExpr)
    if err != nil {
        return err
    }

    // Write the result to the output writer
    if _, err := ch.writer.Write([]byte(prefixExpr)); err != nil {
        return err
    }

    return nil
}
