# Go 1.21 機能デモ

Go 1.21 で追加された主要な機能を実際のコードで解説します。

## ビルドと実行

```bash
go build .    # ビルド
go run .      # 実行
go test ./... # テスト
```

## 各機能の解説

### 1. `log/slog` — 構造化ログ

key-value ペアでログを出力でき、JSON ハンドラへの
切り替えも一行で可能。ログレベルの制御も
`HandlerOptions` で統一的に管理できます。

```go
logger := slog.New(slog.NewTextHandler(os.Stdout,
    &slog.HandlerOptions{Level: slog.LevelInfo},
))
logger.Info("slog demo",
    "feature", "structured logging", "version", "1.21")
```

Go 1.20 以前ではサードパーティの zerolog や zap に
頼る必要がありました。標準ライブラリだけで
構造化ログが書けるようになり、依存関係を減らせます。

---

### 2. `slices` パッケージ — Sort, Contains, Index

独自の比較関数（less 関数）を書く手間が省け、
スライスの探索やソートを 1 行のメソッド呼び出しで
完了できます。

```go
// After (Go 1.21)
slices.Sort(nums)
found := slices.Contains(nums, 4)
```

```go
// Before (Go 1.20)
sort.Ints(nums) // 型ごとにメソッドが分かれていた
// 探索は for ループで自作する必要がありました
```

---

### 3. `maps` パッケージ — Keys, Values

マップのキー一覧や値一覧が欲しいとき、
ボイラープレートな for ループを書かずに済みます。

```go
// After (Go 1.21)
keys := maps.Keys(m)
```

```go
// Before (Go 1.20)
keys := make([]string, 0, len(m))
for k := range m {
    keys = append(keys, k)
}
```

キー一覧を取得するだけのコードが大幅に削減されます。

---

### 4. 組み込み `min` / `max` / `clear`

整数型の最小・最大値を求めるためだけに `float64` への変換
（`math.Min`）をする必要がなくなります。
また、マップの再利用が容易になります。

```go
// After (Go 1.21)
m := min(a, b, c) // 整数でも何でもOK
clear(myMap)      // マップを空にする（メモリ領域は再利用）
```

```go
// Before (Go 1.20)
m := a
if b < m { m = b } // または math.Min(float64(a), float64(b))
// マップのクリアは for k := range m { delete(m, k) }
```

---

## まとめ

| 機能 | パッケージ | 概要 |
| --- | --- | --- |
| 構造化ログ | `log/slog` | key-value ペアでのログ出力 |
| スライス操作 | `slices` | Sort, Contains, Index 等 |
| マップ操作 | `maps` | Keys, Values の取得 |
| 組み込み関数 | （言語機能） | `min`, `max`, `clear` |
