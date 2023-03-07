package main

import (
    "strings"
    "flag"
    "log"
    "os"
    "io"

    lab2 "github.com/nikitosikvn1/SEC-Lab-2"
)

var (
    exprFlag = flag.String("e", "", "Expression to compute")
    fileFlag = flag.String("f", "", "Path to the input file")
    outFlag = flag.String("o", "", "Path to the output file")
)

func main() {
    flag.Parse()

    // Getting input expression
    var reader io.Reader
    if *exprFlag != "" {
        reader = strings.NewReader(*exprFlag)
    } else if *fileFlag != "" {
        inpfile, err := os.Open(*fileFlag)
        if err != nil {
            log.Fatal(err)
        }
        defer inpfile.Close()
        reader = inpfile
    } else {
        log.Fatal("either -e or -f flag must be provided")
    }
    
    // Output of the resulting expression
    var writer io.Writer
    if *outFlag != "" {
        outfile, err := os.Create(*outFlag)
        if err != nil {
            log.Fatal(err)
        }
        defer outfile.Close()
        writer = outfile
    } else {
        writer = os.Stdout
    }
    
    // Call handler
    handler := lab2.NewComputeHandler(reader, writer)
    err := handler.Compute()
    if err != nil {
        log.Fatal(err)
    }

    // TODO: Change this to accept input from the command line arguments as described in the task and
    //       output the results using the ComputeHandler instance.
    //       handler := &lab2.ComputeHandler{
    //           Input: {construct io.Reader according the command line parameters},
    //           Output: {construct io.Writer according the command line parameters},
    //       }
    //       err := handler.Compute()

    // res, _:= lab2.PostfixToPrefix("1 2 3 * + 4 +")
    // fmt.Println(res)
}