# C++23 機能デモ

C++23 で追加された主要な機能を実際のコードで解説します。

## ビルドと実行

```bash
make          # ビルド
make run      # 実行
make test     # テストビルド
make runtest  # テスト実行
```

## 各機能の解説

### 1. 多次元添字演算子 (Multidimensional `operator[]`)

C++23 より、`operator[]`が複数の引数を取れるようになりました。
従来は `m[i][j]` のように連鎖させるしかありませんでしたが、`m[i, j]` と直接書けるようになります。

```cpp
class Matrix2D {
    int data_[3][3]{};
public:
    int& operator[](size_t r, size_t c) { return data_[r][c]; }
};

Matrix2D m;
m[0, 0] = 1;   // C++23: OK
m[1, 2] = 7;   // C++23: OK
```

C++20 以前では `operator[]` は引数 1 個しか取れず、
`m[0][0]` のように連鎖させるか、`operator()(r, c)` を使う必要がありました。

---

### 2. `std::expected<T, E>` — 値かエラーかを返す型

関数の戻り値として「成功値」または「エラー情報」のどちらかを保持する型です。
例外を使わずにエラーハンドリングができます。

```cpp
#include <expected>
#include <string>

std::expected<int, std::string> safe_divide(int a, int b) {
    if (b == 0)
        return std::unexpected("division by zero");  // エラー
    return a / b;                                     // 成功
}

auto r = safe_divide(10, 3);
if (r) {
    printf("結果: %d\n", *r);          // 値にアクセス
} else {
    printf("エラー: %s\n", r.error().c_str());  // エラーにアクセス
}
```

`std::optional` は値の有無だけでしたが、`std::expected` はエラー理由も保持できます。
Rust の `Result<T, E>` に相当する機能です。

---

### 3. `std::optional` モナド操作 — `transform` / `and_then` / `or_else`

C++23 で `std::optional` にモナド的なチェーン操作が追加されました。
値が存在する場合のみ変換や処理を適用し、流れるようにパイプラインを構築できます。

```cpp
std::optional<int> maybe_half(int n) {
    if (n % 2 != 0) return std::nullopt;
    return n / 2;
}

auto result = maybe_half(8)
    .transform([](int n) { return n * 3; })        // 値があれば変換
    .and_then([](int n) -> std::optional<int> {     // 値があれば処理し、optionalを返す
        return n > 0 ? std::optional{n} : std::nullopt;
    })
    .or_else([] { return std::optional{0}; });      // 値がなければ代替値

// maybe_half(8) → 4 → transform(*3) → 12 → and_then(12>0) → 12
// 結果: 12
```

| メソッド | 動作 |
| --- | --- |
| `.transform(f)` | 値があれば `f` を適用して結果をoptionalで返す |
| `.and_then(f)` | 値があれば `f`（optionalを返す関数）を適用する |
| `.or_else(f)` | 値がなければ `f` を呼び出して代替optionalを返す |

---

### 4. `if consteval` — コンパイル時か実行時かで分岐

コンパイル時に評価されているかどうかで処理を切り替えられます。

```cpp
constexpr auto sq = [](int n) {
    if consteval {
        return n * n;     // コンパイル時: 二乗
    } else {
        return n + n;     // 実行時: 倍算
    }
};

static_assert(sq(5) == 25);   // コンパイル時: 5*5 = 25
printf("sq(5) = %d\n", sq(5)); // 実行時: 5+5 = 10
```

C++20 の `if constexpr` が型に基づく分岐だったのに対し、`if consteval` は評価コンテキストに基づく分岐です。
コンパイル時には最適化版、実行時には別の実装を使いたい場合に便利です。

---

### 5. Deducing `this` — 明示的オブジェクトパラメータ

メンバ関数の最初の引数に `this` を明示的に書けるようになりました。
これにより、値・参照・const 修飾ごとのオーバーロードを 1 つの関数にまとめられます。

```cpp
struct Doubler {
    // `this const Doubler&` が明示的オブジェクトパラメータ
    int apply(this const Doubler& self, int n) {
        return n * 2;
    }
};

printf("result: %d\n", Doubler{}.apply(21));  // 42
```

主なメリット：

- const/非 const、左辺値/右辺値のオーバーロードを 1 関数に集約できる
- 再帰的なラムダや CRTP がシンプルに書ける
- `self` という名前を自由に付けられる

---

## まとめ

| 機能 | ヘッダ | 概要 |
| --- | --- | --- |
| 多次元 `operator[]` | （言語機能） | `m[r, c]` で直接アクセス |
| `std::expected` | `<expected>` | 値またはエラーを返す型 |
| optional モナド操作 | `<optional>` | `transform`/`and_then`/`or_else` |
| `if consteval` | （言語機能） | コンパイル時/実行時の分岐 |
| Deducing `this` | （言語機能） | 明示的オブジェクトパラメータ |
