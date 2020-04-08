package main

import (
	"fmt"

	"./mywaifulab"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(entry)
}

func entry() {
	fmt.Println("in entry")
	mywaifulab.GetImage()
	mywaifulab.SendToSlack()
}
