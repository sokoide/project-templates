# Rust 2018 機能デモ

Rust 2018 で追加・改善された主要な機能を実際のコードで解説します。

## ビルドと実行

```bash
make build    # ビルド
make run      # 実行
make test     # テスト実行
```

## 各機能の解説

### 1. `async` / `.await` — 非同期処理のファーストクラスサポート

高並行な非同期処理を、同期コードに近い読みやすい構文で記述でき、スレッドのコンテキストスイッチのオーバーヘッドを最小化します。

```rust
async fn say_hello() {
    println!("hello from async function!");
}

let future = say_hello();
block_on(future); // .await 等で実行
```

---

### 2. Non-Lexical Lifetimes (NLL) — 借用チェッカの高度化

変数の「字面上のスコープ」ではなく「最後に使われた場所」で借用期間を判断するようになり、コードの本質に関わらない借用エラーから解放されます。

```rust
let mut x = 5;
let y = &x;
println!("y is {}", y); // ここで y の借用は終わる
x = 6; // 以前のエディションでは y のスコープが終わるまで x を変更できなかった
```

---

### 3. `dyn` Trait — 動的ディスパッチの明示

トレイトオブジェクトであることを `dyn` キーワードで明記するようになり、静的ディスパッチと動的ディスパッチを明確に区別してコードの意図を伝えやすくなります。

```rust
let animal: Box<dyn Animal> = Box::new(Dog);
animal.speak();
```

---

## まとめ

| 機能 | 概要 |
| --- | --- |
| `async` / `.await` | 言語レベルでの非同期構文サポート |
| NLL | 借用チェッカの緩和による生産性向上 |
| `dyn` | トレイトオブジェクトの明示的なマーク |
| Module System | `extern crate` の廃止などパス指定の簡略化 |
