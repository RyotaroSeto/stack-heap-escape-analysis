package main

import (
	"fmt"
	"testing"
)

func main() {
	fmt.Println(foo())
	fmt.Println(bar())
	loop()
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

func loop() {
	for v := range 3 {
		fmt.Println(&v) // エスケープされていたため、各イテレーションで異なるインスタンスが割り当てられる
		// 0x140000a0030
		// 0x140000a0038
		// 0x140000a0040
	}
	for v := range 3 {
		// 構造体をインターフェース変換してそのメソッドを呼ぶ実装がないため
		println(&v) // エスケープされないため、ループ変数が使いまわされる
		// 0x14000096eb8
		// 0x14000096eb8
		// 0x14000096eb8
	}
}
