# Rust 2015 機能デモ

Rust 2015 (1.0) で確立された根幹となる機能を実際のコードで解説します。
C++ / Go のエキスパートが Rust で最初に直面する概念を網羅しています。

## ビルドと実行

```bash
make build    # ビルド
make run      # 実行
make test     # テスト実行
```

## 各機能の解説

### 1. Ownership & Borrowing — 所有権と借用

ガベージコレクションなしでメモリ安全性を完全に保証する、
Rust の根幹となる仕組みです。

```rust
let s1 = String::from("hello");
let s2 = s1; // Move: 所有権が移動し、s1 は使用不可
// println!("{}", s1); // コンパイルエラー

let s3 = String::from("world");
print_string(&s3); // 借用: 所有権は移動しない
```

C++ のムーブセマンティクスに似ていますが、
ムーブ後の変数は「未規定状態」ではなく**完全に使用不可**になります。
Go には所有権の概念がなく、GC が管理します。

---

### 2. Pattern Matching — パターンマッチング

`match` で値の構造に応じた分岐を書けます。
列挙子の網羅性がコンパイル時に検査されるため、
ハンドリング漏れを防げます。

```rust
let x = Some(5);
match x {
    Some(i) => println!("got Some({})", i),
    None => println!("got None"),
}
```

C++ の `std::visit` や Go の `switch` よりも
はるかに強力で安全です。

---

### 3. Option と Result — null と例外の代替

Rust には `null` も `try-catch` 例外もありません。

- **`Option<T>`**: 値がある（`Some`）かない（`None`）かを型で表現
- **`Result<T, E>`**: 成功（`Ok`）か失敗（`Err`）かを型で表現

```rust
fn divide(a: f64, b: f64) -> Result<f64, String> {
    if b == 0.0 {
        // format! を使った動的なエラーメッセージ
        Err(format!("Cannot divide {} by zero", a))
    } else {
        Ok(a / b)
    }
}

let result = divide(10.0, 3.0)?;
```

Go の `if err != nil` に相応するのが `?` 演算子です。
C++ の `std::optional` / `std::expected` に相当しますが、
`?` による伝播が言語レベルでサポートされています。

---

### 4. Trait — トレイト（インターフェース）

共通の振る舞いを定義する仕組みです。
Go のインターフェースや C++ の概念（Concepts）に似ていますが、
静的ディスパッチが基本で、必要に応じて動的ディスパッチ（`dyn`）も使えます。

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

Go と違い、トレイトの実装を明示的に宣言します。

---

### 5. `String` vs `&str` — 文字列の2種類

Rust を学ぶ際に**最も多くの人が混乱するポイント**です。

| 型 | 所有権 | 用途 | Go / C++ との対応 |
| --- | --- | --- | --- |
| `String` | ヒープ確保、所有権あり | 文字列の生成・変更 | Go の `string`、C++ の `std::string` |
| `&str` | 借用（参照） | 文字列の読み取り専用 | Go の `string`（読み取り）、C++ の `std::string_view` |

```rust
let owned: String = String::from("hello");  // ヒープ上
let borrowed: &str = &owned;                 // 借用
let literal: &str = "world";                 // 静的領域への参照

fn greet(name: &str) { // 引数は &str で統一
    println!("Hello, {}!", name);
}
greet(&owned);  // String も &str も渡せる
```

**実務のコツ**: 関数の引数は `&str` で受けると、
`String` も `&str` も渡せるため最も柔軟です。

---

### 6. `Copy` vs `Clone` — コピーの2種類

C++ はデフォルトでコピーされ、Go も値型はコピーされますが、
Rust は**型ごとにコピーの動作が明示的に異なります**。

- **`Copy`**: ビットコピーだけで安全な型（`i32`, `f64`, `bool` 等）
- **`Clone`**: 明示的な `.clone()` 呼び出しが必要な型（`String`, `Vec<T>` 等）

```rust
let a = 42;
let b = a;    // Copy: i32 は自動コピー、a も使える
println!("{}", a); // OK

let s1 = String::from("hello");
let s2 = s1;  // Move: String は Copy ではない
// println!("{}", s1); // エラー: 所有権が移動済み

let s3 = String::from("world");
let s4 = s3.clone(); // Clone: 明示的な複製
println!("{} {}", s3, s4); // 両方使える
```

**C++ との違い**: C++ では `std::string s2 = s1;` はコピーですが、
Rust ではムーブになります。コピーしたい場合は `.clone()` が必要です。

---

### 7. `Vec` / 配列 / スライス — コレクションの3層構造

Go の `[]T` に相当するものが、所有権の有無で 3 種類あります。

| 型 | 所有権 | サイズ | Go / C++ との対応 |
| --- | --- | --- | --- |
| `Vec<T>` | 所有権あり（ヒープ） | 可変長 | Go の `[]T`、C++ の `std::vector` |
| `[T; N]` | 所有権あり（スタック） | 固定長 | C++ の `std::array<T, N>` |
| `&[T]` | 借用（参照） | 任意 | Go の `[]T`（読み取り）、C++ の `std::span` |

```rust
let mut v: Vec<i32> = Vec::new();
v.push(1);
v.push(2);

let arr: [i32; 3] = [1, 2, 3];

fn sum(slice: &[i32]) -> i32 { // スライスで受けると何でも渡せる
    slice.iter().sum()
}
sum(&v);   // OK: Vec からスライスへ
sum(&arr); // OK: 配列からスライスへ
```

**実務のコツ**: 関数の引数は `&[T]` で受けると、
`Vec` も配列も渡せるため最も柔軟です。

---

### 8. `Drop` トレイト — RAII とデストラクタ

C++ のデストラクタに相当します。
値がスコープを抜けるタイミングで自動的にリソースを解放します。

```rust
struct FileHandle {
    name: String,
}

impl Drop for FileHandle {
    fn drop(&mut self) {
        println!("Dropping: {}", self.name);
    }
}

{
    let f = FileHandle { name: String::from("data.txt") };
    // f を使う
} // ここで自動的に Drop::drop が呼ばれる
```

C++ の RAII と同じ考え方ですが、
Rust では `Drop` を実装した型に対して明示的な `drop` 呼び出しはできません
（所有権のルールにより、コンパイラが自動で管理します）。

---

### 9. `Box` / `Rc` / `Arc` — スマートポインタ

C++ のスマートポインタや Go のポインタに対応する型です。

| 型 | 所有権 | Go / C++ との対応 |
| --- | --- | --- |
| `Box<T>` | 単一所有権 | C++ の `std::unique_ptr` |
| `Rc<T>` | 共有所有権（単一スレッド） | C++ の `std::shared_ptr` |
| `Arc<T>` | 共有所有権（マルチスレッド安全） | C++ の `std::shared_ptr`（スレッド安全） |

```rust
// ヒープ確保（サイズが不明なトレイトオブジェクト等に使用）
let b: Box<i32> = Box::new(42);

// 参照カウント（単一スレッド）
use std::rc::Rc;
let a = Rc::new(vec![1, 2, 3]);
let b = Rc::clone(&a); // 参照カウントが増える、データは共有
```

Go にはスマートポインタの概念がなく、
ポインタはすべて GC が管理します。

---

### 10. イテレータとクロージャ

C++ の `<algorithm>` + ラムダや、
Go の `for range` に対応する Rust の書き方です。

```rust
let nums = vec![1, 2, 3, 4, 5];

// map + collect（C++ の std::transform 相当）
let doubled: Vec<i32> = nums.iter().map(|x| x * 2).collect();

// filter（C++ の std::copy_if 相当）
let evens: Vec<&i32> = nums.iter().filter(|x| *x % 2 == 0).collect();

// fold（C++ の std::accumulate 相当）
let sum: i32 = nums.iter().fold(0, |acc, x| acc + x);
```

クロージャの構文は `|args| body` で、
C++ の `[capture](args) { body }` や
Go の `func(args) { body }` に相当します。

---

### 11. モジュールシステム — `mod` / `pub` / `use`

Go の `package` / C++ の `#include` に対応するコード分割の仕組みです。

```rust
mod math {
    pub fn add(a: i32, b: i32) -> i32 { // pub で外部に公開
        a + b
    }
    fn secret() {} // 非 pub はモジュール内限定
}

use math::add; // use でインポート
```

Go と違い、可視性は `pub` キーワードで明示します
（Go は大文字始まりで公開）。C++ には言語レベルのモジュールが
C++20 で導入されましたが、Rust の方が洗練されています。

---

### 12. `derive` マクロ — ボイラープレートの自動生成

`#[derive(...)]` で一般的なトレイト実装を自動生成できます。

```rust
#[derive(Debug, Clone, PartialEq)]
struct Point {
    x: f64,
    y: f64,
}

let p1 = Point { x: 1.0, y: 2.0 };
let p2 = p1.clone();
println!("{:?}", p1);     // Debug による整形出力
assert_eq!(p1, p2);       // PartialEq による比較
```

C++ にはコンパイラ生成のデフォルト操作がありますが、
Rust の `derive` はマクロベースでより明示的です。

---

### 13. Lifetimes — ライフタイム

参照が有効な期間をコンパイラに伝える仕組みです。
ほとんどの場合は推論されますが、関数の引数と戻り値が
複数の参照に関わる場合は明示が必要です。

```rust
// 戻り値の参照がどの引数と同じ寿命かを明示
fn longest<'a>(x: &'a str, y: &'a str) -> &'a str {
    if x.len() > y.len() { x } else { y }
}
```

C++ にも参照の寿命はありますが、コンパイラは検査しません（ダングリング参照の危険）。
Go には参照の寿命の概念がなく、GC が管理します。
Rust はこれを**コンパイル時に検査**する点が独特です。

---

## まとめ

| 機能 | 概要 |
| --- | --- |
| Ownership & Borrowing | コンパイル時の静的メモリ管理 |
| Pattern Matching | 強力かつ網羅的な条件分岐 |
| Option / Result | null と例外の代替、`?` 演算子 |
| Trait | 静的ディスパッチを基本としたポリモーフィズム |
| `String` vs `&str` | 所有権付き文字列と借用文字列の使い分け |
| `Copy` vs `Clone` | 暗黙コピーと明示的複製の区別 |
| `Vec` / 配列 / スライス | 所有権に基づくコレクションの3層構造 |
| `Drop` | RAII パターンによる自動リソース解放 |
| `Box` / `Rc` / `Arc` | 所有権モデルに基づくスマートポインタ |
| イテレータ / クロージャ | 関数型スタイルのデータ処理 |
| モジュールシステム | `pub` / `use` による可視性制御 |
| `derive` マクロ | ボイラープレートの自動生成 |
| Lifetimes | 参照の有効期間のコンパイル時検査 |

## C++ / Go との対応まとめ

| Rust | C++ | Go |
| --- | --- | --- |
| `String` | `std::string` | `string` |
| `&str` | `std::string_view` | `string`（読み取り） |
| `Vec<T>` | `std::vector<T>` | `[]T` |
| `&[T]` | `std::span<T>` | `[]T`（読み取り） |
| `Box<T>` | `std::unique_ptr<T>` | `*T`（GC 管理） |
| `Rc<T>` / `Arc<T>` | `std::shared_ptr<T>` | なし（GC 管理） |
| `Option<T>` | `std::optional<T>` | `*T`（nil チェック） |
| `Result<T, E>` | `std::expected<T, E>` | `(T, error)` |
| `trait` | Concept / Interface | `interface{}` |
| `match` | `std::visit` | `switch` |
| `?` 演算子 | なし | `if err != nil` |
| `Drop` | デストラクタ | `defer` |
