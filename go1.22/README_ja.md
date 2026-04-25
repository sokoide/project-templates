# Go 1.22 機能デモ

Go 1.22 で追加された主要な機能を実際のコードで解説します。

## ビルドと実行

```bash
go build .    # ビルド
go run .      # 実行
go test ./... # テスト
```

## 各機能の解説

### 1. ループ変数のスコープ変更

for ループの変数が反復ごとに新しく作成される
ようになり、ゴルーチンでよく起きるキャプチャバグが
消滅します。

```go
for i := 0; i < 3; i++ {
    go func() {
        results = append(results, i)
        // Go 1.22: 各ゴルーチンが固有の i を持つ
    }()
}
```

Go 1.21 以前ではループ変数は全体で共有されるため、
ゴルーチンが実行される時点で `i` が 3 に達していました。
回避するには `i := i` や引数渡しが必要でした。

```go
// Go 1.21 以前の回避策
for i := 0; i < 3; i++ {
    i := i // シャドウイングで固有の変数を作る
    go func() {
        results = append(results, i)
    }()
}
```

---

### 2. HTTP ルーティング — メソッドマッチングとパスパラメータ

外部ルーター（chi, gorilla/mux 等）を導入せずとも、
標準ライブラリだけで RESTful な API を簡潔かつ安全に
構築できます。

```go
// After (Go 1.22)
mux.HandleFunc("GET /items/{id}", func(w http.ResponseWriter, r *http.Request) {
    id := r.PathValue("id")
    fmt.Fprintf(w, "item: %s", id)
})
```

```go
// Before (Go 1.21)
// 1. パスを strings.Split で解析して ID を抽出する
// 2. r.Method == "GET" を自前でチェックする
// 3. または外部ライブラリを依存関係に追加する
```

メソッドチェック（405）やワイルドカード抽出が
標準化されました。

---

## まとめ

| 機能 | パッケージ | 概要 |
| --- | --- | --- |
| ループ変数スコープ | （言語機能） | 反復ごとに固有の変数を生成 |
| HTTP ルーティング | `net/http` | メソッドマッチング、パスパラメータ |
