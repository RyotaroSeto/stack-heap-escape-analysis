package main

import (
	"fmt"
	"testing"
)

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

func Escape() *int {
	x := 1 // NOTE: Escape関数return後もポインタが参照されるためエスケープされる。10倍近くパフォーマンスが落ちる
	return &x
}

func NoEscape() int {
	y := 1
	return y
}

func BenchmarkEscape(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Escape()
	}
}
