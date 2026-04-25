# Rust 2024 機能デモ

Rust 2024 で導入された主要な変更や、将来的な進化に向けた基盤を解説します。

## ビルドと実行

```bash
make build    # ビルド
make run      # 実行
make test     # テスト実行
make check    # 静的解析 (fmt & clippy)
```

## 各機能の解説

### 1. `impl Trait` capture rules — キャプチャルールの簡素化

`impl Trait` を戻り値に持つ関数において、インスコープにあるすべての
ライフタイムをデフォルトでキャプチャするようになります。
以前のエディションで必要だった `+ 'a` などの明示が不要になります。

```rust
fn create_printer<'a>(msg: &'a str) -> impl Printer {
    SimplePrinter { msg } // 2024: 'a が暗黙的にキャプチャされる
}
```

### 2. `gen` keyword — ジェネレータへの布石

`gen` キーワードが予約語となり、将来的に `Iterator` や `Stream` をネイティブに生成する `gen { yield ... }` 構文の導入が予定されています。

---

## まとめ

| 特徴 | 内容 |
| --- | --- |
| Lifetime Capture | `impl Trait` におけるライフタイム推論の改善 |
| Reserved Keywords | `gen` キーワードの予約による将来的な拡張性 |
| Library Stability | 標準ライブラリのさらなる安定化と洗練 |
