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

「値がある」または「値がない」のどちらかを表現できる型です。
エラーコードやセンチネル値を使わずに、安全に「見つからなかった」ことを伝えられます。

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

複数の型のどれか一つを保持できる型です。
従来の `union` と違い、どの型が入っているかを型安全に追跡できます。

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

従来の `union` はどのメンバが有効かをプログラマが手動で管理する必要があり、誤用すると未定義動作になりました。`std::variant` は型情報を内部で追跡するため安全です。

---

### 3. `std::string_view` — 文字列の非所有参照

`std::string` や文字列リテラルのデータをコピーせずに参照する軽量な型です。
関数の引数として文字列を受け取る際、所有権のない参照として効率よく渡せます。

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

C++17 以前では、`const std::string&` を使うと暗黙のメモリ確保が発生し、
`const char*` を使うとサイズ計算が毎回走るというトレードオフがありました。
`std::string_view` はポインタと長さのペアを保持するため、
どちらのデメリットも回避できます。

---

### 4. 構造化束縛 (Structured bindings)

構造体、`std::pair`、`std::tuple`、配列などの要素を個別の変数に一度に束縛できます。

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

#### `std::tuple` の場合

```cpp
auto [x, y, z] = std::make_tuple(1, 2.5, "three");
// x == 1, y == 2.5, z == "three"
```

C++17 以前では、`std::get<0>(tuple)` や `pair.first` / `pair.second` で個別にアクセスする必要がありました。

---

### 5. Fold expressions — 可変長引数の畳み込み

可変長テンプレート引数パックに対して二項演算子を適用できます。
再帰的なテンプレート定義なしに、パラメータパックを簡潔に処理できます。

```cpp
template <typename... Args>
auto fold_sum(Args... args) {
    return (args + ...);  // 左fold: ((args[0] + args[1]) + args[2]) + ...
}

printf("fold_sum(1..5)=%d\n", fold_sum(1, 2, 3, 4, 5));            // 15
printf("fold_sum(1.1,2.2,3.3)=%.1f\n", fold_sum(1.1, 2.2, 3.3));  // 6.6
```

C++17 以前では、可変長引数に対する演算を行うために再帰的なテンプレート関数を定義する必要がありました。Fold expressions により 1 行で書けるようになりました。

---

### 6. `if constexpr` — コンパイル時条件分岐

コンパイル時に型の条件に基づいてコードの有無を切り替えられます。
条件を満たさない分岐はコンパイルされないため、テンプレートの特殊化が不要になります。

```cpp
template <typename T>
std::string stringify(T val) {
    if constexpr (std::is_integral_v<T>) {
        return "int:" + std::to_string(val);
    } else {
        return "float:" + std::to_string(val);
    }
}

printf("stringify(42)=%s\n", stringify(42).c_str());       // int:42
printf("stringify(3.14)=%s\n", stringify(3.14).c_str());   // float:3.140000
```

C++17 以前では、型ごとにテンプレート特殊化やオーバーロードを用意する必要がありました。`if constexpr` を使えば 1 つの関数テンプレート内で分岐できます。

---

### 7. CTAD (Class Template Argument Deduction) — クラステンプレート引数推論

クラステンプレートのテンプレート引数をコンストラクタの引数から自動推論します。
明示的にテンプレート引数を書く必要がなくなりました。

```cpp
std::pair p{1, 2};          // std::pair<int, int> と推論される
std::optional opt{99};       // std::optional<int> と推論される

printf("CTAD pair: (%d,%d)\n", p.first, p.second);
printf("CTAD optional: %d\n", *opt);
```

C++17 以前では `std::pair<int, int> p{1, 2}` のようにテンプレート引数を明示するか、
`std::make_pair(1, 2)` などのヘルパー関数を使う必要がありました。

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
