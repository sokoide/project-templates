# Go 1.26 機能デモ

Go 1.26 で追加された主要な機能をコードで解説します。

## ビルドと実行

```bash
go build .    # ビルド
go run .      # 実行
go test ./... # テスト
```

## 各機能の解説

### 1. `new` 式初期化 — ポインタに値を指定して生成

JSON パラメータやオプション設定で頻出する
「値へのポインタ」を一時変数なしで一行で書けます。

```go
// After (Go 1.26)
p := Person{
    Age: new(yearsSince(born)),
}
```

```go
// Before (Go 1.25)
age := yearsSince(born)
p := Person{
    Age: &age, // 直接 &yearsSince(born) とは書けなかった
}
```

---

### 2. 自己参照型制約 (Self-referential type constraints)

「自分と同じ型を引数に取る」ような再帰的な型制約を、
ハックなしで自然に記述できるようになります。

```go
// After (Go 1.26)
type Adder[A Adder[A]] interface {
    Add(A) A
}
```

```go
// Before (Go 1.25)
// 同様の定義をしようとすると
// "invalid recursive type" エラーが発生し、
// ジェネリクスを用いない interface 定義などで
// 妥協する必要がありました。
```

---

### 3. `errors.AsType` — ジェネリック版 `errors.As`

ターゲット変数の宣言を分離する必要がなくなり、
`if` 文の条件式の中で完結するため変数スコープを
最小限に抑えられます。

```go
// After (Go 1.26)
if ne, ok := errors.AsType[*NotFoundError](err); ok {
    fmt.Println(ne.Name)
}
```

```go
// Before (Go 1.25)
var ne *NotFoundError
if errors.As(err, &ne) {
    fmt.Println(ne.Name)
}
```

---

### 4. `reflect` イテレータ — 構造体フィールドの反復

リフレクションを用いたメタプログラミングにおいて、
泥臭いインデックスループが消え `range` で直感的に
スキャンできるようになります。

```go
// After (Go 1.26)
for field, val := range reflect.ValueOf(p).Fields() {
    fmt.Printf("%s=%v\n", field.Name, val.Int())
}
```

```go
// Before (Go 1.25)
t := reflect.TypeOf(p)
for i := 0; i < t.NumField(); i++ {
    fmt.Println(t.Field(i).Name)
}
```

---

### 5. `slog.NewMultiHandler` — 複数ログハンドラの統合

「標準出力とファイルの双方にログを出したい」といった
定番の要件が、外部ライブラリや自前実装なしで
1 行で解決します。

```go
// After (Go 1.26)
handler := slog.NewMultiHandler(stdoutHandler, fileHandler)
```

```go
// Before (Go 1.25)
// 複数のハンドラをラップして各メソッドを転送する
// 構造体を自前で実装するか、サードパーティ製の
// マルチハンドラを利用する必要がありました。
```

---

### 6. `bytes.Buffer.Peek` — 読み込み位置を進めずに覗き見

バッファの内容を消費せずに先頭を確認するために、
わざわざ `bufio.Reader` でラップする手間がなくなります。

```go
// After (Go 1.26)
peeked, _ := buf.Peek(5) // 消費せずに中身を確認
```

```go
// Before (Go 1.25)
// bytes.Buffer には Peek がなかったため、
// 1. bufio.NewReader(&buf) でラップする
// 2. buf.Bytes() でスライスを直接触る
//    （書き込みがあると無効になるリスクあり）
// といった工夫が必要でした。
```

---

### 7. Green Tea GC — 新ガベージコレクタのデフォルト化

コードを 1 行も変えることなくコンパイルし直すだけで、
実行時の GC オーバーヘッドが最大 40% 削減されます。

```text
// Go 1.26 では「何もしない」のが最善です。
// ランタイムが自動的に最新のアルゴリズム
// （Green Tea GC）を使用します。
```

| 特徴 | 詳細 |
| --- | --- |
| **改善幅** | 10-40% のオーバーヘッド削減 |
| **仕組み** | 小オブジェクトのスキャン効率化とスケーラビリティ向上 |
| **ハードウェア最適化** | Ice Lake / Zen 4 以降でベクトル命令により加速 |

Go 1.5 の並行化以来、最も大きな GC の進化と言えます。

---

## まとめ

| 機能 | パッケージ | 概要 |
| --- | --- | --- |
| `new` 式初期化 | （言語機能） | `new(expr)` で値付きポインタを生成 |
| 自己参照型制約 | （言語機能） | 再帰的な型制約を自然に記述 |
| `errors.AsType` | `errors` | ジェネリック版 `As`、型安全なエラー抽出 |
| `reflect` イテレータ | `reflect` | `Fields`/`Methods` で range 反復 |
| `NewMultiHandler` | `log/slog` | 複数ハンドラの統合 |
| `Buffer.Peek` | `bytes` | 読み込み位置を進めずに覗き見 |
| Green Tea GC | （ランタイム） | 新 GC デフォルト化、オーバーヘッド削減 |
