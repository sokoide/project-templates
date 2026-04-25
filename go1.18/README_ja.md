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

スライス操作などの共通ロジックを、型ごとにコピペすることなく、コンパイル時の型安全性を保ったまま共通化できます。

```go
// After
func Map[T any, U any](s []T, f func(T) U) []U { ... }

ints := []int{1, 2, 3}
strs := Map(ints, func(i int) string { return strconv.Itoa(i) })
```

```go
// Before
// 1. 各型ごとに同様のループ処理を個別に書く
// 2. interface{} で受け取り、実行時にリフレクションや型アサーションを行う（低速かつ危険）
```

---

### 2. ジェネリック型（Stack[T]）

独自のコンテナ型（スタック、キュー、ツリーなど）を実装する際、利用時の型アサーションを完全に排除し、安全にデータを扱えます。

```go
// After
type Stack[T any] struct { items []T }
s := Stack[int]{}
s.Push(10)
val := s.Pop() // キャスト不要で int 型として取り出せる
```

```go
// Before
// []interface{} を使い、Pop のたびに v.(int) などの型アサーションが必要でした。
```

---

### 3. 型制約（Ordered）— ユニオン要素

「数値型なら何でも受け取る」といった制約を宣言的に記述でき、ジェネリック関数内で比較演算子（`>` や `<`）を安全に使用できるようになります。

```go
// After
type Ordered interface {
    ~int | ~float64 | ~string
}
func Max[T Ordered](a, b T) T {
    if a > b { return a } // Ordered 制約により比較演算が可能
    return b
}
```

```go
// Before
// 比較演算を行う汎用関数を、複数の型に対して安全に記述する手段がありませんでした。
```

---

## まとめ

| 機能 | パッケージ | 概要 |
| --- | --- | --- |
| ジェネリック関数 | （言語機能） | 型パラメータによる汎用関数 |
| ジェネリック型 | （言語機能） | 型パラメータ付き構造体 |
| 型制約 | （言語機能） | ユニオン要素で許可型を列挙 |
