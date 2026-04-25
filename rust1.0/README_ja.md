# Rust 1.0 機能デモ

Rust 1.0 で確立された根幹となる機能を実際のコードで解説します。

## ビルドと実行

```bash
make build    # ビルド
make run      # 実行
make test     # テスト実行
```

## 各機能の解説

### 1. Ownership & Borrowing — 所有権と借用

メモリ管理の責任をコンパイラに移譲し、ガベージコレクション（GC）なしでメモリ安全性を完全に保証します。

```rust
let s = String::from("hello");
print_string(&s); // 借用（データの所有権は渡さない）
let _s2 = s;      // ムーブ（所有権が移動し、sは使用不能になる）
```

---

### 2. Pattern Matching — パターンマッチング

複雑な条件分岐を網羅的かつ安全に記述でき、`None` やエラーのハンドリング漏れをコンパイル時に防ぎます。

```rust
let x = Some(5);
match x {
    Some(i) => println!("got Some({})", i),
    None => println!("got None"),
}
```

---

### 3. Trait — トレイト

データとその振る舞い（インターフェース）を分離して定義でき、抽象度の高いコードを型安全に再利用できます。

```rust
trait HasArea {
    fn area(&self) -> f64;
}

impl HasArea for Circle {
    fn area(&self) -> f64 {
        std::f64::consts::PI * self.radius * self.radius
    }
}
```

---

## まとめ

| 機能 | 概要 |
| --- | --- |
| Ownership | コンパイル時の静的メモリ管理 |
| Borrowing | 不変・可変参照によるデータ共有の制御 |
| Pattern Matching | 強力かつ網羅的な条件分岐 |
| Trait | 静的ディスパッチを基本としたポリモーフィズム |
