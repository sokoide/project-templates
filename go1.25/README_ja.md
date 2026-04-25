# Go 1.25 機能デモ

Go 1.25 で追加された主要な機能をコードで解説します。

## ビルドと実行

```bash
go build .    # ビルド
go run .      # 実行
go test ./... # テスト
```

## 各機能の解説

### 1. `testing/synctest` — 決定的な並行テスト

ゴルーチンの終了を待つための `time.Sleep` が不要になり、
テストの実行時間が短縮され Flaky テストを根絶できます。

```go
// After (Go 1.25)
synctest.Run(func() {
    go longRunningTask()
    synctest.Wait() // 確実に完了を待機
})
```

```go
// Before (Go 1.24)
// time.Sleep(100 * time.Millisecond)
// // これで動くはず...（不安定なテストの原因）
```

---

### 2. 並行コレクトパターン — ゴルーチン + チャネル

複数のゴルーチンから結果をチャネルで集約する定番
パターンです。`WaitGroup` で完了を待ち、メイン
ゴルーチンで結果を収集します。

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
    fmt.Println(v) // 0, 10, 20（順序は不定）
}
```

---

### 3. パイプラインパターン — チャネルで接続

複数の処理ステージをチャネルで直列接続します。
各ステージは独立したゴルーチンで並行実行されます。

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

// Stage 2: 値を 2 倍
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

`select` と `time.After` を組み合わせてタイムアウトを
実装します。指定時間内に応答がなければ
タイムアウト側のケースが実行されます。

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
