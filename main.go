package main

import (
	"fmt"
	//"github.com/Luwade/jenkins/go/src/github.com/AsynkronIT/goconsole"
)

func main()  {
	sum_integer := Sum()
	fmt.Println("The sum of two integers are: ", sum_integer)
}

func Sum() int {
	sum := 1 + 1
	return sum
}