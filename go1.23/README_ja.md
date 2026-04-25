# Go 1.23 機能デモ

Go 1.23 で追加された主要な機能をコードで解説します。

## ビルドと実行

```bash
go build .    # ビルド
go run .      # 実行
go test ./... # テスト
```

## 各機能の解説

### 1. `iter.Seq` — 単一値カスタムイテレータ

独自のデータ構造や生成ロジックを `range` に対応させ、
利用者側が内部構造を意識せずにループを回せます。

```go
// After (Go 1.23)
func Countdown(n int) iter.Seq[int] { ... }
for v := range Countdown(5) { // 5, 4, 3, 2, 1
    fmt.Println(v)
}
```

```go
// Before (Go 1.22)
// 1. スライスを生成して返す（全データをメモリに積む）
// 2. Next() メソッドを持つ構造体を作り、
// 2. Next() メソッドなどを持つ構造体を作り、
// 2. Next() メソッドなどを持つ構造体を作り、
//    for ステートメントで手動で回す
```

---

### 2. `iter.Seq2` — キー・値カスタムイテレータ

二つの値を返すイテレータでインデックス付き反復などを
自然に記述できます。

```go
func Enumerate[T any](s []T) iter.Seq2[int, T] {
    return func(yield func(int, T) bool) {
        for i, v := range s {
            if !yield(i, v) {
                return
            }
        }
    }
}

for i, w := range Enumerate([]string{"a", "b"}) {
    fmt.Printf("%d:%s ", i, w) // 0:a 1:b
}
```

---

### 3. Fibonacci イテレータ — 無限イテレータと break

無限に値を生成するイテレータを定義でき、`break` で
途中終了時にクリーンアップ処理を記述できます。

```go
func Fibonacci() iter.Seq[int] {
    return func(yield func(int) bool) {
        a, b := 0, 1
        for {
            if !yield(a) {
                return
            }
            a, b = b, a+b
        }
    }
}

count := 0
for v := range Fibonacci() {
    fmt.Println(v)
    count++
    if count >= 8 {
        break
    }
}
```

---

### 4. 標準ライブラリイテレータ — `slices.All`, `maps.All`

スライスやマップをイテレータとして戻り値にでき、
呼び出し側は `range` で一貫して処理できます。

```go
// After (Go 1.23)
func GetItems() iter.Seq2[int, string] {
    return slices.All(items)
}
for i, v := range GetItems() { ... }
```

```go
// Before (Go 1.22)
// スライスそのものを返すしかなく、内部データの露出や
// 巨大データのコピーコストが発生していました。
```

---

## まとめ

| 機能 | パッケージ | 概要 |
| --- | --- | --- |
| `iter.Seq` | `iter` | 単一値カスタムイテレータ |
| `iter.Seq2` | `iter` | キー・値カスタムイテレータ |
| Fibonacci | `iter` | 無限イテレータと break による中断 |
| `slices.All` | `slices` | スライスのインデックス・値イテレータ |
| `maps.All` | `maps` | マップのキー・値イテレータ |
