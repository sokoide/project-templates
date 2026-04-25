# C++17 機能デモ

C++17 で追加された主要な機能を実際のコードで解説します。

## ビルドと実行

```bash
make          # ビルド
make run      # 実行
make test     # テストビルド
make runtest  # テスト実行
```

## 各機能の解説

### 1. `std::optional` — 値の有無を表す型

エラーコードやセンチネル値を使わずに、安全に「値がない」ことを呼び出し側に伝えられます。

```cpp
#include <optional>

std::optional<int> find_index(const std::vector<int> &v, int target) {
    auto it = std::find(v.begin(), v.end(), target);
    if (it != v.end())
        return static_cast<int>(it - v.begin());
    return std::nullopt;  // 値なし
}

auto idx = find_index(v, 30);
if (idx)
    printf("index=%d\n", *idx);   // 値にアクセス
else
    printf("not found\n");         // 値なし
```

C++17 以前では、戻り値として `-1` や `nullptr` などのセンチネル値を使う必要があり、呼び出し側での判定が曖昧になりがちでした。

---

### 2. `std::variant` — 型安全な共用体 (Type-safe union)

どの型が入っているかをランタイムで安全に追跡でき、不適切な型へのアクセスを未然に防ぎます。

```cpp
#include <variant>

using Value = std::variant<int, double, std::string>;

Value vi = 42;
Value vd = 3.14;
Value vs = std::string("hello");

// std::visitで型に応じた処理を書ける
std::string value_type_name(const Value &v) {
    return std::visit(
        [](auto &&arg) -> std::string {
            using T = std::decay_t<decltype(arg)>;
            if constexpr (std::is_same_v<T, int>)
                return "int";
            else if constexpr (std::is_same_v<T, double>)
                return "double";
            else if constexpr (std::is_same_v<T, std::string>)
                return "string";
            else
                return "unknown";
        },
        v);
}
```

従来の `union` はどのメンバが有効かをプログラマが手動で管理する必要があり、誤用すると未定義動作になりました。

---

### 3. `std::string_view` — 文字列の非所有参照

文字列をコピーせずに参照でき、メモリ確保のオーバーヘッドなしに効率よく文字列操作を行えます。

```cpp
#include <string_view>

size_t count_words(std::string_view sv) {
    size_t count = 0;
    bool in_word = false;
    for (char c : sv) {
        if (c == ' ' || c == '\t') {
            in_word = false;
        } else if (!in_word) {
            ++count;
            in_word = true;
        }
    }
    return count;
}

printf("word count: %zu\n", count_words("hello world foo bar"));  // 4
```

C++17 以前では、`const std::string&` を使うと暗黙のメモリ確保が発生するリスクがありました。

---

### 4. 構造化束縛 (Structured bindings)

構造体やタプルの要素を個別の変数に一度に展開でき、冗長なアクセス用コードを排除できます。

#### 構造体の場合

```cpp
struct DivResult {
    int quotient;
    int remainder;
};

DivResult int_div(int a, int b) { return {a / b, a % b}; }

auto [q, r] = int_div(17, 5);
// q == 3, r == 2
```

#### `std::map` のイテレーション

```cpp
std::map<std::string, int> scores;
scores["alice"] = 90;
scores["bob"] = 85;
for (auto &[name, score] : scores) {
    printf("  %s: %d\n", name.c_str(), score);
}
```

C++17 以前では、`pair.first` / `pair.second` などで個別にアクセスする必要がありました。

---

### 5. Fold expressions — 可変長引数の畳み込み

再帰的なテンプレート定義なしに、パラメータパックに対する演算を 1 行で簡潔に記述できます。

```cpp
template <typename... Args>
auto fold_sum(Args... args) {
    return (args + ...);  // 左fold: ((args[0] + args[1]) + args[2]) + ...
}

printf("fold_sum(1..5)=%d\n", fold_sum(1, 2, 3, 4, 5));            // 15
```

C++17 以前では、可変長引数を処理するために複雑な再帰テンプレートが必要でした。

---

### 6. `if constexpr` — コンパイル時条件分岐

コンパイル時に条件に合わない分岐を完全に排除でき、テンプレートの特殊化を減らしてコードをシンプルに保てます。

```cpp
template <typename T>
std::string stringify(T val) {
    if constexpr (std::is_integral_v<T>) {
        return "int:" + std::to_string(val);
    } else {
        return "float:" + std::to_string(val);
    }
}
```

---

### 7. CTAD (Class Template Argument Deduction) — クラステンプレート引数推論

コンストラクタの引数から型を自動推論させることで、煩わしいテンプレート引数の明示を省略できます。

```cpp
std::pair p{1, 2};          // std::pair<int, int> と推論される
std::optional opt{99};       // std::optional<int> と推論される
```

C++17 以前では `std::make_pair(1, 2)` などのヘルパー関数が必要でした。

---

## まとめ

| 機能 | ヘッダ | 概要 |
| --- | --- | --- |
| `std::optional` | `<optional>` | 値の有無を表現する型 |
| `std::variant` | `<variant>` | 型安全な共用体 |
| `std::string_view` | `<string_view>` | 文字列の非所有軽量参照 |
| 構造化束縛 | （言語機能） | 構造体/pair/tupleの要素を個別変数に束縛 |
| Fold expressions | （言語機能） | 可変長引数パックの畳み込み |
| `if constexpr` | （言語機能） | コンパイル時の型に基づく条件分岐 |
| CTAD | （言語機能） | クラステンプレート引数の自動推論 |
