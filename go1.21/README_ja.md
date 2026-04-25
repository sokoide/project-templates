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

標準ライブラリだけで key-value ペアによる構造化ログを出力でき、サードパーティライブラリへの依存を減らしつつ柔軟なログ管理が可能になります。

```go
logger := slog.New(slog.NewTextHandler(os.Stdout,
    &slog.HandlerOptions{Level: slog.LevelInfo},
))
logger.Info("slog demo",
    "feature", "structured logging", "version", "1.21")
```

Go 1.20 以前では zerolog や zap に頼る必要がありましたが、標準化によりエコシステム全体の相互運用性が向上しました。

---

### 2. `slices` パッケージ — Sort, Contains, Index

型ごとに異なるソート関数を呼び出す手間が省け、スライスの探索や操作をジェネリクスによって型安全かつ 1 行で完結できます。

```go
// After
slices.Sort(nums)
found := slices.Contains(nums, 4)
```

```go
// Before
sort.Ints(nums) // 型ごとにメソッドが分かれていた
// 探索は for ループで自作する必要がありました
```

---

### 3. `maps` パッケージ — Keys, Values

マップのキー一覧や値一覧を取得するだけの冗長な for ループが不要になり、ボイラープレートコードを大幅に削減できます。

```go
// After
keys := maps.Keys(m)
```

```go
// Before
keys := make([]string, 0, len(m))
for k := range m {
    keys = append(keys, k)
}
```

---

### 4. 組み込み `min` / `max` / `clear`

最小・最大値の計算を、型変換（`math.Min`）や `if` 文による自作なしで直感的に記述でき、マップのクリア処理も 1 行で効率的に行えます。

```go
// After
m := min(a, b, c) // 任意の比較可能型に対応
clear(myMap)      // マップを効率的に空にする
```

```go
// Before
m := a
if b < m { m = b }
// マップのクリアは for ループで delete を繰り返す必要がありました
```

---

## まとめ

| 機能 | パッケージ | 概要 |
| --- | --- | --- |
| 構造化ログ | `log/slog` | key-value ペアでのログ出力 |
| スライス操作 | `slices` | Sort, Contains, Index 等 |
| マップ操作 | `maps` | Keys, Values の取得 |
| 組み込み関数 | （言語機能） | `min`, `max`, `clear` |
