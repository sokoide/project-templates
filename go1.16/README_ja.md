# Go 1.16 機能デモ

Go 1.16 で追加された主要な機能を実際のコードで解説します。

## ビルドと実行

```bash
go build .    # ビルド
go run .      # 実行
go test ./... # テスト
```

## 各機能の解説

### 1. `//go:embed` — 文字列への埋め込み

設定ファイルや SQL クエリなどのテキストファイルを、
外部ファイルとして管理しつつ、
実行バイナリに直接組み込めます。

```go
// After (Go 1.16)
//go:embed static/hello.txt
var helloStr string
```

```go
// Before (Go 1.15)
// 1. ioutil.ReadFile で実行時に読み込む（パス管理が面倒）
// 2. 文字列リテラルとして .go ファイルにコピペする（同期が大変）
```

---

### 2. `//go:embed` — バイトスライスへの埋め込み

画像やバイナリデータを、アセット管理のフローを
崩さずにそのまま型安全に扱えます。

```go
// After (Go 1.16)
//go:embed static/icon.png
var iconBytes []byte
```

```go
// Before (Go 1.15)
// go-bindata などの外部ツールを使って .go ファイルに変換する
// 必要がありました。
```

---

### 3. `//go:embed` — `embed.FS` への埋め込み

シングルバイナリで Web サーバーを構築でき、
フロントエンドの静的アセットを
バイナリ一つにパッキングして配布できます。

```go
// After (Go 1.16)
//go:embed static
var staticFS embed.FS

data, _ := fs.ReadFile(staticFS, "static/hello.txt")
```

```go
// Before (Go 1.15)
// ディレクトリ構造を維持したまま埋め込むには、
// 外部ツールで巨大なソースコードを生成する必要がありました。
```

---

## まとめ

| 機能 | パッケージ | 概要 |
| --- | --- | --- |
| 文字列埋め込み | `embed` | ファイルを `string` に埋め込み |
| バイト埋め込み | `embed` | ファイルを `[]byte` に埋め込み |
| FS 埋め込み | `embed` + `io/fs` | ディレクトリを仮想 FS に埋め込み |
