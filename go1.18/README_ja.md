# Go 1.18 機能デモ

Go 1.18 で追加された主要な機能を実際のコードで解説します。

## ビルドと実行

```bash
go build .    # ビルド
go run .      # 実行
go test ./... # テスト
```

## 各機能の解説

### 1. ジェネリック関数（Map, Filter, Reduce）

スライス操作などの共通ロジックを、型ごとにコピペ
することなく型安全に共通化できます。

```go
// After (Go 1.18)
func Map[T any, U any](s []T, f func(T) U) []U { ... }

ints := []int{1, 2, 3}
strs := Map(ints, func(i int) string { return strconv.Itoa(i) })
```

```go
// Before (Go 1.17)
// 1. 各型（[]int, []string）ごとに同様のループ処理を書く
// 2. interface{} で受け取り、実行時にリフレクションや
//    型アサーションを行う（遅くて危険）
```

---

### 2. ジェネリック型（Stack[T]）

独自のコンテナ型（スタック、キュー、ツリーなど）を
実装する際、利用時に型アサーションを繰り返す必要が
なくなります。

```go
// After (Go 1.18)
type Stack[T any] struct { items []T }
s := Stack[int]{}
s.Push(10)
val := s.Pop() // int 型として取り出せる
```

```go
// Before (Go 1.17)
// items []interface{} を使い、Pop のたびに v.(int) と
// 型アサーションが必要でした。
```

---

### 3. 型制約（Ordered）— ユニオン要素

「数値型なら何でも受け取る」といった制約を
インターフェースで表現でき、演算子（`>` や `+`）が
使える型を安全に制限できます。

```go
// After (Go 1.18)
type Ordered interface {
    ~int | ~float64 | ~string // ...他
}
func Max[T Ordered](a, b T) T {
    if a > b { return a } // Ordered なので比較演算が可能
    return b
}
```

```go
// Before (Go 1.17)
// 比較演算を行う汎用関数を型安全に書く手段がありませんでした。
```

---

## まとめ

| 機能 | パッケージ | 概要 |
| --- | --- | --- |
| ジェネリック関数 | （言語機能） | 型パラメータによる汎用関数 |
| ジェネリック型 | （言語機能） | 型パラメータ付き構造体 |
| 型制約 | （言語機能） | ユニオン要素で許可型を列挙 |
