package main

import (
    "bufio"
    "flag"
    "fmt"
    "log"
    "os"

    lab2 "github.com/nikitosikvn1/SEC-Lab-2"
)

var (
    inputExpression = flag.String("e", "", "Expression to compute")
    inputFile = flag.String("f", "", "Path to the input file")
    outputFile = flag.String("o", "", "Path to the output file")

    exp string
    // TODO: Add other flags support for input and output configuration.
)

func main() {
    flag.Parse()
    
    // If a value was passed for the -e flag, then we read the expression
    if *inputExpression != "" {
        exp = *inputExpression

    // If a path was passed for the -f flag, then we read the contents of the file
    } else if *inputFile != "" {
        fileInfo, err := os.Stat(*inputFile)
        if err != nil {
            log.Fatal("error: ", err)
        }

        if fileInfo.IsDir() {
            log.Fatal(*inputFile, " is a directory, not a file")
        }

        file, err := os.Open(*inputFile)
        if err != nil {
            log.Fatal("error: ", err)
        }
        
        defer file.Close()

        scanner := bufio.NewScanner(file)
        scanner.Scan()
        exp = scanner.Text()

    // If -e, -f did not receive a value, then terminate the program
    } else {
        log.Fatal("error. No input was passed")
    }

    fmt.Println(exp)

    // TODO: Change this to accept input from the command line arguments as described in the task and
    //       output the results using the ComputeHandler instance.
    //       handler := &lab2.ComputeHandler{
    //           Input: {construct io.Reader according the command line parameters},
    //           Output: {construct io.Writer according the command line parameters},
    //       }
    //       err := handler.Compute()

    res, _:= lab2.PostfixToPrefix("1 2 3 * + 4 +")
    fmt.Println(res)
}