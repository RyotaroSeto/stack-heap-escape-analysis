package main

import "fmt"

func main() {
	fmt.Println(foo())
	fmt.Println(bar())
}

func foo() *int {
	i := 42
	return &i // iはエスケープしてヒープに割り当てられます
}

func bar() int {
	j := 100
	return j // jはエスケープせずスタックに割り当てられます
}
