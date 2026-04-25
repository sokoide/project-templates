# C++20 機能デモ

C++20 で追加された主要な機能を実際のコードで解説します。

## ビルドと実行

```bash
make          # ビルド
make run      # 実行
make test     # テストビルド
make runtest  # テスト実行
```

## 各機能の解説

### 1. Concepts（コンセプト）

テンプレート引数に対する制約を宣言的に記述できるようになりました。
型パラメータが満たすべき条件を明示でき、エラーメッセージも分かりやすくなります。

```cpp
#include <concepts>

template <std::integral T> T generic_add(T a, T b);
template <std::floating_point T> T generic_add(T a, T b);

generic_add(1, 2);       // OK: int は std::integral
generic_add(1.5, 2.5);   // OK: double は std::floating_point
```

C++17 以前では、SFINAE と `std::enable_if` を使って同様の制約を記述していましたが、構文が複雑でエラーメッセージも読みにくいものでした。

---

### 2. Three-way Comparison（三方比較・宇宙船演算子 `<=>`）

`<=>` 演算子（宇宙船演算子）により、`==`、`!=`、`<`、`<=`、`>`、`>=` の 6 種類の比較演算子を一度に定義できます。

```cpp
#include <compare>

struct Point {
    int x, y;
    auto operator<=>(const Point &) const = default;
};

Point a{1, 2}, b{3, 4}, c{1, 2};
auto cmp = (a <=> b);                                     // less
printf("a==c: %d\n", a == c);                             // 1 (true)
printf("a<=>b: %s\n", std::is_lt(cmp) ? "less" : ...);    // less
```

`= default` と書くだけでメンバごとの辞書式比較が自動生成されます。
C++17 以前では、各比較演算子を個別に実装する必要がありました。

---

### 3. `std::span`

配列やコンテナの一部を所有権を移転せずに参照する軽量ビューです。
生ポインタとサイズのペアを安全に置き換えられます。

```cpp
#include <span>

int span_sum(std::span<const int> s) {
    int sum = 0;
    for (auto v : s) sum += v;
    return sum;
}

int arr[] = {10, 20, 30, 40, 50};
std::vector<int> v = {1, 2, 3};

span_sum(arr);     // OK: C配列から暗黙変換
span_sum(v);       // OK: std::vectorから暗黙変換
```

C++17 以前では、`int*` と `size_t` を別々に受け取るか、テンプレートでコンテナ全体を受け取る必要がありました。

---

### 4. Designated Initializers（指示付き初期化子）

構造体のメンバを名前で指定して初期化できます。
初期化順序は宣言順である必要がありますが、省略されたメンバはデフォルト値で初期化されます。

```cpp
struct Config {
    int width = 0;
    int height = 0;
    bool fullscreen = false;
};

Config make_default_config() {
    return {.width = 800, .height = 600, .fullscreen = false};
}

Config custom{.width = 1920, .height = 1080, .fullscreen = true};
```

C 言語にはありましたが、C++では C++20 でようやくサポートされました。
C++17 以前ではコンストラクタや集約初期化 `{1920, 1080, true}` を使う必要があり、順序を覚える必要がありました。

---

### 5. `consteval`

`consteval` で宣言された関数は、必ずコンパイル時に評価されます。
実行時に呼び出そうとするとコンパイルエラーになります。

```cpp
consteval int square(int n) { return n * n; }

static_assert(square(5) == 25);   // コンパイル時評価: OK
printf("square(5)=%d\n", square(5));  // 定数式として評価
```

C++17 の `constexpr` は「コンパイル時にも評価できる」でしたが、`consteval` は「コンパイル時にしか評価できない」というより強い保証を持ちます。

---

### 6. Ranges（レンジライブラリ）

要素の列に対する操作をパイプライン形式で組み立てられるようになりました。

#### レンジアルゴリズム

従来の `<algorithm>` のレンジ版で、イテレータのペアではなくコンテナを直接渡せます。

```cpp
#include <ranges>
#include <algorithm>

std::vector<int> filter_even(const std::vector<int> &input) {
    std::vector<int> result;
    std::ranges::copy_if(input, std::back_inserter(result),
                         [](int n) { return n % 2 == 0; });
    return result;
}

std::vector<int> transform_double(const std::vector<int> &input) {
    std::vector<int> result;
    std::ranges::transform(input, std::back_inserter(result),
                           [](int n) { return n * 2; });
    return result;
}
```

#### ビューパイプライン

遅延評価のビューを `|` で繋いでパイプラインを構築できます。

```cpp
namespace rv = std::views;

std::vector<int> nums = {1, 2, 3, 4, 5, 6};

// 偶数を抽出して二乗
auto result = nums | rv::filter([](int n) { return n % 2 == 0; })
                   | rv::transform([](int n) { return n * n; });
// result: 4, 16, 36

// 連番を生成して先頭3個を取得
auto first3 = rv::iota(10, 20) | rv::take(3);
// first3: 10, 11, 12
```

C++17 以前では、一時変数を複数作るか、`<algorithm>` に `begin()`/`end()` を渡す必要がありました。

---

## まとめ

| 機能 | ヘッダ | 概要 |
| --- | --- | --- |
| Concepts | `<concepts>` | テンプレート引数への制約 |
| 三方比較 `<=>` | `<compare>` | 比較演算子の自動生成 |
| `std::span` | `<span>` | コンテナの非所有ビュー |
| 指示付き初期化子 | （言語機能） | メンバ名指定の初期化 |
| `consteval` | （言語機能） | コンパイル時評価の強制 |
| Ranges | `<ranges>` | パイプライン形式の要素操作 |
