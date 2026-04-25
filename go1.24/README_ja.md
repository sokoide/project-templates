# Go 1.24 機能デモ

Go 1.24 で追加された主要な機能をコードで解説します。

## ビルドと実行

```bash
go build .    # ビルド
go run .      # 実行
go test ./... # テスト
```

## 各機能の解説

### 1. ジェネリック定義型 — メソッド付き型パラメータ型

型パラメータを持つ定義型にメソッドを定義できるように
なり、ジェネリックコレクションを自然に実装できます。

```go
type Set[T comparable] map[T]struct{}

func NewSet[T comparable]() Set[T] {
    return make(Set[T])
}

s := NewSet[string]()
s.Add("go")
s.Contains("go") // true
```

Go 1.23 以前ではジェネリック定義型のメソッド定義に制限が
あり、型パラメータを含む定義型でコンパイルエラーに
なるケースがありました。

---

### 2. `os.Root` — ファイルシステムルート制約

外部パスが意図しないディレクトリ（`/etc/passwd` など）
を指していないかを気にする必要がなくなり、OS レベルの
安全なサンドボックスを利用できます。

```go
// After (Go 1.24)
root, _ := os.OpenRoot("/user/uploaded")
f, err := root.Open(userInput) // パストラバーサルを自動防御
```

```go
// Before (Go 1.23)
// filepath.Join("/user/uploaded", userInput) したあと、
// 結果がプレフィックスを持つか自前でチェックが必要
```

---

### 3. `go.mod` tool ディレクティブ — ツール依存管理

開発用 CLI ツール（`stringer`, `mockgen` 等）のバージョンを
プロジェクト全体で統一でき、`tools.go` のハックから
卒業できます。

```text
// After (Go 1.24)
// go.mod
tool github.com/golang/mock/mockgen
```

```go
// Before (Go 1.23)
// tools.go というファイルを作り、ビルドタグ
// // +build tools を付け、インポート文を並べて
// バージョンを固定する面倒な作業が必要でした。
```

---

### 4. `filepath.Localize` — 安全なパス検証

ファイルパスを検証しローカルパスとして安全か確認
します。パストラバーサルを含むパスはエラーです。
パストラバーサルを含むパスはエラーになります。
パストラバーサルを含むパスはエラーになります。

```go
safe := filepath.Localize("subdir/file.txt")
// => "subdir/file.txt"

// filepath.Localize("../../etc/passwd")
// => エラー（パストラバーサル）
```

Go 1.23 以前では自前で `filepath.Clean` と
`strings.HasPrefix` 等を使った検証が必要でした。

---

## まとめ

| 機能 | パッケージ | 概要 |
| --- | --- | --- |
| Generic type alias | （言語機能） | 型エイリアスに型パラメータを指定可能 |
| `os.Root` | `os` | ディレクトリをルートとした安全な FS アクセス |
| tool ディレクティブ | `go.mod` | ツール依存を go.mod で一元管理 |
| `filepath.Localize` | `path/filepath` | ローカルパスの安全性を検証 |
