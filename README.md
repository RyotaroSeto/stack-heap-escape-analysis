# stack-heap-escape-analysis

## エスケープ解析の確認方法
```sh
go build -gcflags '-m' main.go

# command-line-arguments
./main.go:10:6: can inline foo
./main.go:15:6: can inline bar
./main.go:6:17: inlining call to foo
./main.go:6:13: inlining call to fmt.Println
./main.go:7:17: inlining call to bar
./main.go:7:13: inlining call to fmt.Println
./main.go:6:17: moved to heap: i
./main.go:6:13: ... argument does not escape
./main.go:7:13: ... argument does not escape
./main.go:7:17: ~r0 escapes to heap
./main.go:11:2: moved to heap: i
```

## それぞれ
- Go で「変数がエスケープする」とは、変数へのポインタがスタック領域ではなくヒープ領域に置かれる(退避する)こと
- スタックは各 goroutine が管理するメモリ領域で、レキシカルスコープ内の静的な(ライフサイクルが予測できる)データが格納される
  - FIFO で、参照されるスコープも明 確なためハンドリングしやすいという特徴がある
- ヒープはレキシカルスコープ外のデータや、サイズが大きすぎる・予見できないデータが主に格納される
  - ある変数が宣言元の関数がreturnした後もポインタを参照されるケースは、エスケープが起きる代表的な例
  - ヒープ割り当ては、変数のライフサイクルをGCで管理する必要があることから、スタック割り当てと比べてパフォーマンスが落ちる傾向がある。
