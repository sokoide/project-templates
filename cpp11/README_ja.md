# C++11 機能デモ

C++11 で追加された主要な機能を実際のコードで解説します。

## ビルドと実行

```bash
make          # ビルド
make run      # 実行
make test     # テストビルド
make runtest  # テスト実行
```

> **注意:** GoogleTest 1.17 は C++17+ を要求するため、テストコード
> (`test_main.cpp`) は `-std=c++17` でコンパイルされます。
> メインコードは `-std=c++11` でコンパイルされています。

## 各機能の解説

### 1. `auto` — 型推論

**エンジニアが喜ぶポイント:** 複雑なイテレータ型や長いテンプレート名を何度も書く苦行から解放されます。

変数の型を初期化子から自動的に推論します。

```cpp
// After (C++11)
auto x = 42;                        // int
auto name = std::string("hello");   // std::string
```

```cpp
// Before (C++98)
int x = 42;
std::string name = std::string("hello");
```

---

### 2. Range-based `for` — 範囲for文

**エンジニアが喜ぶポイント:** ループの境界条件ミス（オフバイワンエラー）が構造的に発生しなくなります。

配列やコンテナの全要素を簡潔に反復できます。

```cpp
// After (C++11)
std::vector<int> v = {1, 2, 3, 4, 5};
for (auto n : v)
    printf(" %d", n);
```

```cpp
// Before (C++98)
for (std::vector<int>::iterator it = v.begin(); it != v.end(); ++it)
    printf(" %d", *it);
```

---

### 3. Lambda — ラムダ式（キャプチャ付き）

その場で無名関数オブジェクトを定義できます。 `[...]` で周囲の変数をキャプチャできます。

```cpp
// キャプチャなし
auto greet = [](const char *s) { printf("lambda: %s\n", s); };
greet("hello");

// キャプチャ付き: [capture] で外の変数を取り込む
int capture = 10;
auto add = [capture](int n) { return n + capture; };
printf("lambda capture: %d\n", add(5));  // 15
```

`std::function` と組み合わせて、関数を返す関数も作れます。

```cpp
// sub.cpp
std::function<int(int)> make_adder(int base) {
    return [base](int x) { return base + x; };
}

// main.cpp
auto add5 = make_adder(5);
printf("adder(5)(10)=%d\n", add5(10));  // 15
```

C++11 以前では、関数オブジェクト（ファンクタ）用の構造体を定義する必要がありました。

```cpp
// C++11以前
struct Adder {
    int base;
    Adder(int b) : base(b) {}
    int operator()(int x) const { return base + x; }
};
```

---

### 4. `enum class` — スコープ付き列挙型

従来の `enum` は暗黙に `int` に変換され、名前が周囲のスコープに漏れていました。`enum class` は型安全でスコープが限定されます。

```cpp
// sub.h
enum class Color { Red, Green, Blue };

// sub.cpp
std::string color_name(Color c) {
    switch (c) {
    case Color::Red:   return "Red";
    case Color::Green: return "Green";
    case Color::Blue:  return "Blue";
    }
    return "Unknown";
}

// main.cpp
Color c = Color::Green;
printf("enum class: %s\n", color_name(c).c_str());
```

C++11 以前の `enum` では、名前衝突や暗黙の型変換が問題でした。

```cpp
// C++11以前: 名前がグローバルに漏れる、暗黙にintに変換される
enum Color { Red, Green, Blue };
int n = Red;  // OK（意図しない変換）
```

---

### 5. `constexpr` — コンパイル時定数式

関数や変数をコンパイル時に評価できることを示します。テンプレートメタプログラミングの代わりに、通常の関数構文で定数計算が書けます。

```cpp
constexpr int factorial(int n) { return n <= 1 ? 1 : n * factorial(n - 1); }

// コンパイル時に評価される
printf("constexpr factorial(5)=%d\n", factorial(5));  // 120
```

C++11 以前では、コンパイル時計算にテンプレートの再帰を使う必要がありました。

```cpp
// C++11以前
template<int N> struct Factorial {
    static const int value = N * Factorial<N-1>::value;
};
template<> struct Factorial<1> { static const int value = 1; };
```

---

### 6. `std::array` — 固定長配列コンテナ

C スタイル配列のラッパーで、STL コンテナのインターフェース（イテレータ、`size()` など）を持ちます。オーバーヘッドはゼロです。

```cpp
std::array<int, 5> arr = {{10, 20, 30, 40, 50}};

int array_sum(const std::array<int, 5> &arr) {
    int sum = 0;
    for (auto v : arr)
        sum += v;
    return sum;
}
```

C++11 以前では、C スタイル配列を使うか、`boost::array` に頼る必要がありました。

```cpp
// C++11以前
int arr[] = {10, 20, 30, 40, 50};
// size()やイテレータは使えない
```

---

### 7. Smart Pointers — スマートポインタ (`unique_ptr`, `shared_ptr`)

手動の `new`/`delete` に代わる自動メモリ管理です。

- **`std::unique_ptr`**: 単一所有権。コピー不可、ムーブのみ。

  ```cpp
  std::unique_ptr<int> make_unique_val(int v) {
      return std::unique_ptr<int>(new int(v));
  }

  auto up = make_unique_val(42);
  printf("unique_ptr: %d\n", *up);
  ```

- **`std::shared_ptr`**: 共有所有権。参照カウントで自動的に解放。

  ```cpp
  std::shared_ptr<std::string> make_shared_str(const std::string &s) {
      return std::make_shared<std::string>(s);
  }

  auto sp = make_shared_str("smart");
  printf("shared_ptr: %s use_count=%ld\n", sp->c_str(), sp.use_count());
  ```

C++11 以前では、`new`/`delete` を手動で管理するか、`boost::shared_ptr` を使う必要がありました。

```cpp
// C++11以前: メモリリークの危険
int* p = new int(42);
// delete p; を忘れるとリーク
```

---

### 8. `std::function` — 汎用関数ラッパー

任意の呼び出し可能オブジェクト（関数ポインタ、ラムダ、ファンクタ）を統一的に扱える型です。

```cpp
#include <functional>

std::function<int(int)> make_adder(int base) {
    return [base](int x) { return base + x; };
}

auto add5 = make_adder(5);
printf("adder(5)(10)=%d\n", add5(10));  // 15
```

C++11 以前では、関数ポインタやファンクタをテンプレートで受け取るしかありませんでした。

---

### 9. Move Semantics — ムーブセマンティクス

リソースの「所有権移動」により、不要なコピーをなくします。一時オブジェクトからリソースを「盗む」ことができます。

```cpp
class Buffer {
    int *data_;
    size_t size_;
public:
    Buffer(size_t sz);
    ~Buffer();
    Buffer(Buffer &&other) noexcept;           // ムーブコンストラクタ
    Buffer &operator=(Buffer &&other) noexcept; // ムーブ代入
    Buffer(const Buffer &) = delete;           // コピー禁止
    Buffer &operator=(const Buffer &) = delete;
    size_t size() const;
};

Buffer b1(100);
printf("b1 size=%zu\n", b1.size());        // 100
Buffer b2 = std::move(b1);                 // 所有権をb2へ移動
printf("after move: b2 size=%zu\n", b2.size());  // 100
```

C++11 以前では、所有権の移動を表現できず、高コストなディープコピーが発生していました。

---

### 10. Initializer Lists — 初期化子リスト

`{...}` 構文でコンテナを簡潔に初期化できます。

```cpp
// std::initializer_listを受け取る関数
std::vector<int> make_sorted(std::initializer_list<int> vals) {
    std::vector<int> v(vals);
    std::sort(v.begin(), v.end());
    return v;
}

// 呼び出し側
auto sorted = make_sorted({5, 3, 1, 4, 2});
// sorted: {1, 2, 3, 4, 5}
```

C++11 以前では、配列を宣言してから個別に代入するか、`push_back` を繰り返す必要がありました。

```cpp
// C++11以前
std::vector<int> v;
v.push_back(5);
v.push_back(3);
v.push_back(1);
v.push_back(4);
v.push_back(2);
std::sort(v.begin(), v.end());
```

---

## まとめ

| 機能 | ヘッダ | 概要 |
| --- | --- | --- |
| `auto` | （言語機能） | 変数の型を初期化子から自動推論 |
| Range-based `for` | （言語機能） | コンテナ・配列の簡潔な反復 |
| Lambda | （言語機能） | 無名関数オブジェクト、キャプチャ付き |
| `enum class` | （言語機能） | 型安全なスコープ付き列挙型 |
| `constexpr` | （言語機能） | コンパイル時定数計算 |
| `std::array` | `<array>` | 固定長配列のSTLコンテナラッパー |
| `unique_ptr` / `shared_ptr` | `<memory>` | 自動メモリ管理のスマートポインタ |
| `std::function` | `<functional>` | 汎用関数ラッパー |
| Move Semantics | （言語機能） | リソース所有権の移動による高速化 |
| Initializer Lists | `<initializer_list>` | `{...}` による統一的初期化構文 |
