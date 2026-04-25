# Go 1.26 機能デモ

Go 1.26 で追加された主要な機能を実際のコードで解説します。

## ビルドと実行

```bash
go build .    # ビルド
go run .      # 実行
go test ./... # テスト
```

## 各機能の解説

### 1. `new` 式初期化 — ポインタに値を指定して生成

組み込み関数 `new` の引数に式を渡せるようになりました。
これにより、ポインタフィールドの初期化が一行で書けます。

```go
type Person struct {
    Name string `json:"name"`
    Age  *int   `json:"age"`
}

born := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
p := Person{
    Name: "Alice",
    Age:  new(yearsSince(born)), // Go 1.26: 式の結果へのポインタを生成
}
```

Go 1.25 以前では `new(int)` でゼロ値のポインタを作ってから代入するか、
ヘルパー関数を定義する必要がありました。

```go
// Go 1.25 以前
age := yearsSince(born)
p := Person{Name: "Alice", Age: &age}
```

JSON のオプショナルフィールドや protocol buffers のポインタフィールドで
特に便利です。

---

### 2. 自己参照型制約 (Self-referential type constraints)

ジェネリクスの型制約で自分自身を参照できるようになりました。
「自分と同じ型を引数に取る」ようなインターフェースを自然に表現できます。

```go
type Adder[A Adder[A]] interface {
    Add(A) A
}

type Int int

func (a Int) Add(b Int) Int { return a + b }

func algo[A Adder[A]](x, y A) A {
    return x.Add(y)
}

algo(Int(3), Int(4)) // => 7
```

Go 1.25 以前では `Adder[A Adder[A]]` のような自己参照がコンパイルエラーに
なっていました。実用的な例として、順序付きコレクションの要素型に
`Compare(T) int` を要求する制約などが書きやすくなります。

---

### 3. `errors.AsType` — ジェネリック版 `errors.As`

エラー値から特定の型を型安全に取り出せるジェネリック関数です。
`errors.As` の型アサーション版で、より簡潔に書けます。

```go
type NotFoundError struct {
    Name string
}

func (e *NotFoundError) Error() string { return "not found: " + e.Name }

err := findItem("widget")
if ne, ok := errors.AsType[*NotFoundError](err); ok {
    fmt.Println(ne.Name) // "widget"
}
```

Go 1.25 以前の `errors.As` ではターゲット変数を `any` 型で宣言し、
関数にポインタを渡す必要がありました。

```go
// Go 1.25 以前
var ne *NotFoundError
if errors.As(err, &ne) {
    fmt.Println(ne.Name)
}
```

`AsType` は型引数で目的の型を明示するため、
間違った型を渡すリスクがありません。

---

### 4. `reflect` イテレータ — 構造体フィールド・メソッドの反復

`reflect` パッケージにイテレータメソッドが追加されました。
`range` で構造体のフィールドやメソッドを簡潔に列挙できます。

```go
type Point struct{ X, Y int }

p := Point{X: 10, Y: 20}

// 構造体の型情報を反復
for field := range reflect.TypeOf(p).Fields() {
    fmt.Println(field.Name) // "X", "Y"
}

// 構造体の値を反復（フィールド情報と値のペア）
for field, val := range reflect.ValueOf(p).Fields() {
    fmt.Printf("%s=%v\n", field.Name, val.Int())
}
```

Go 1.25 以前では `NumField()` と `Field(i)` を使った
インデックスループを書く必要がありました。

```go
// Go 1.25 以前
t := reflect.TypeOf(p)
for i := 0; i < t.NumField(); i++ {
    fmt.Println(t.Field(i).Name)
}
```

---

### 5. `slog.NewMultiHandler` — 複数ログハンドラの統合

複数の `slog.Handler` を一つにまとめる `MultiHandler` が追加されました。
ログをファイルと標準出力に同時出力するような用途で便利です。

```go
textHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})
handler := slog.NewMultiHandler(textHandler)
logger := slog.New(handler)
logger.Info("multi-handler demo", "key", "value")
```

Go 1.25 以前では自前でマルチハンドラを実装するか、
サードパーティライブラリに頼る必要がありました。

---

### 6. `bytes.Buffer.Peek` — 読み込み位置を進めずに覗き見

`bytes.Buffer` に `Peek` メソッドが追加されました。
バッファの先頭 N バイトを、読み込み位置を進めずに取得できます。

```go
var buf bytes.Buffer
buf.WriteString("hello world")

peeked, _ := buf.Peek(5)
fmt.Printf("peeked: %s\n", peeked) // "hello"

fmt.Println(buf.Len()) // 11（Peek は消費しない）
```

Go 1.25 以前では `bufio.Reader` を使うか、
バッファの `Bytes()` を使って手動でスライスを取得する必要がありました。
`bytes.Buffer` で直接 `Peek` できることで、
パーサーやプロトコル実装が書きやすくなります。

---

### 7. Green Tea GC — 新ガベージコレクタのデフォルト化

Go 1.25 で実験導入された Green Tea GC がデフォルト有効化されました。
小オブジェクトのマーク・スキャンにおける局所性と CPU スケーラビリティが改善され、
GC オーバーヘッドが 10–40% 削減されています。

```text
# 無効化する場合（Go 1.27 でオプトアウト設定は削除予定）
GOEXPERIMENT=nogreenteagc go build .
```

コードからの直接の呼び出しはありませんが、
すべての Go 1.26 プログラムで自動的に適用されます。
Intel Ice Lake / AMD Zen 4 以降の CPU では
ベクトル命令を活用したスキャンによりさらに 10% 程度の改善が見込まれます。

Go の GC は 1.5 の並行化以来、段階的に改善されてきましたが、
Green Tea GC はその中でも最も大きな飛躍です。

---

## まとめ

| 機能 | パッケージ | 概要 |
| --- | --- | --- |
| `new` 式初期化 | （言語機能） | `new(expr)` で値付きポインタを生成 |
| 自己参照型制約 | （言語機能） | ジェネリクスで自分自身を参照する制約 |
| `errors.AsType` | `errors` | ジェネリック版 `As`、型安全なエラー抽出 |
| `reflect` イテレータ | `reflect` | `Fields`/`Methods` で range 反復 |
| `NewMultiHandler` | `log/slog` | 複数ハンドラの統合 |
| `Buffer.Peek` | `bytes` | 読み込み位置を進めずにバッファを覗き見 |
| Green Tea GC | （ランタイム） | 新 GC デフォルト化、オーバーヘッド 10–40% 削減 |
