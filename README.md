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
