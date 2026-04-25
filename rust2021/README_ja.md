# Rust 2021 機能デモ

Rust 2021 で追加・洗練された主要な機能を実際のコードで解説します。

## ビルドと実行

```bash
make build    # ビルド
make run      # 実行
make test     # テスト実行
make check    # 静的解析 (fmt & clippy)
```

## 各機能の解説

### 1. Disjoint capture in closures — クロージャの分離キャプチャ

クロージャが構造体全体ではなく「実際に使用するフィールド」のみを個別にキャプチャするようになり、借用の競合によるコンパイルエラーを劇的に減らします。

```rust
let mut p = Player { name: "Alice".into(), score: 100 };
let c = || println!("{}", p.name); // name のみ借用

p.score += 10; // 2018 ままでは p 全体が借用されていたため、ここでエラーだった
c();
```

---

### 2. `IntoIterator` for arrays — 配列の直接反復

配列を直接 `for` ループに渡して所有権ベースの反復ができるようになり、`.iter()` や参照を明示する手間がなくなります。

```rust
let arr = [1, 2, 3];
for x in arr { // 2018 では &arr や arr.iter() が必要だった
    println!("{}", x);
}
```

---

### 3. New Prelude — 標準ライブラリの利便性向上

`TryInto` などの頻出トレイトが標準でインポートされるようになり、明示的な `use` 文を書かずにモダンな変換処理が記述できます。

```rust
// TryInto が Prelude に追加されたため、個別の use なしで使える
let n: i64 = 10;
let n_i32: i32 = n.try_into()?; // unwrap() ではなく ? で伝播
```

---

## まとめ

| 機能 | 概要 |
| --- | --- |
| Disjoint Capture | フィールド単位のキャプチャによる借用緩和 |
| Array IntoIterator | 配列に対する直感的な for ループ |
| New Prelude | `TryInto`, `TryFrom`, `FromIterator` の標準化 |
| Cargo Resolver v2 | 依存関係の解決アルゴリズムの最適化 |
| Error Handling | テンプレート全体での `?` 演算子の推奨 |
