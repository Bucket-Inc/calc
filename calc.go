package main

import (
	"fmt"
	"github.com/Knetic/govaluate"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: calc <expression>")
		return
	}

	expression := os.Args[1]
	expr, err := govaluate.NewEvaluableExpression(expression)
	if err != nil {
		fmt.Println("Error creating evaluable expression:", err)
		return
	}

	result, err := expr.Evaluate(nil)
	if err != nil {
		fmt.Println("Error evaluating expression:", err)
		return
	}
	fmt.Println("Result:", result)
}

