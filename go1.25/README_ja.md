# Go 1.25 機能デモ

Go 1.25 で追加された主要な機能を実際のコードで解説します。

## ビルドと実行

```bash
go build .    # ビルド
go run .      # 実行
go test ./... # テスト
```

## 各機能の解説

### 1. `testing/synctest` — 決定的な並行テスト

ゴルーチンの完了を確実に待機できるため、不安定な `time.Sleep` を排除し、Flaky テストを根絶できます。

```go
// After
synctest.Run(func() {
    go longRunningTask()
    synctest.Wait() // 確実に完了を待機
})
```

```go
// Before
// time.Sleep(100 * time.Millisecond)
// // これで動くはず...（不安定なテストの原因）
```

---

### 2. 並行コレクトパターン — ゴルーチン + チャネル

複数の並行処理の結果を効率よく 1 箇所に集約し、処理の高速化と安全なデータ収集を両立できます。

```go
var wg sync.WaitGroup
results := make(chan int, 3)

for i := 0; i < 3; i++ {
    wg.Add(1)
    go func(n int) {
        defer wg.Done()
        results <- n * 10
    }(i)
}

go func() {
    wg.Wait()
    close(results)
}()

for v := range results {
    fmt.Println(v) // 0, 10, 20
}
```

---

### 3. パイプラインパターン — チャネルで接続

複雑な処理を独立したステージに分割してチャネルで接続でき、並行性とコードの再利用性を高められます。

```go
in := make(chan int)
out := make(chan int)

// Stage 1: 生成
go func() {
    for i := 1; i <= 3; i++ {
        in <- i
    }
    close(in)
}()

// Stage 2: 加工
go func() {
    for v := range in {
        out <- v * 2
    }
    close(out)
}()

for v := range out {
    fmt.Println(v) // 2, 4, 6
}
```

---

### 4. タイムアウト処理 — select + time.After

応答のない処理を安全に切り捨てることができ、システム全体のデッドロックやリソース枯渇を防止できます。

```go
ch := make(chan string, 1)
ch <- "quick response"

select {
case msg := <-ch:
    fmt.Println(msg)
case <-time.After(time.Second):
    fmt.Println("timed out")
}
```

---

## まとめ

| 機能 | パッケージ | 概要 |
| --- | --- | --- |
| `synctest.Run` | `testing/synctest` | 決定的な並行テスト |
| `synctest.Wait` | `testing/synctest` | ゴルーチンブロックまで待機 |
| 並行コレクト | `sync` | WaitGroup + チャネルで結果集約 |
| パイプライン | `sync` | チャネルでステージを直列接続 |
| タイムアウト | `time` | select + After でタイムアウト |
