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

行列や画像データへのアクセスを、数学的な記法 `m[r, c]` で記述でき、可読性が向上します。

```cpp
// After (C++23)
m[0, 0] = 1; 
```

```cpp
// Before (C++20)
m[0][0] = 1;   // operator[] の連鎖が必要
m(0, 0) = 1;   // または operator() で代用していた
```

---

### 2. `std::expected<T, E>` — 値かエラーかを返す型

例外を投げずに「失敗の理由」を型安全に返せ、呼び出し側でのエラーハンドリングを簡潔かつ確実にします。

```cpp
// After (C++23)
std::expected<int, std::string> result = safe_divide(10, 0);
if (!result) {
    printf("Error: %s\n", result.error().c_str());
}
```

```cpp
// Before (C++20)
// 1. 例外を投げる
// 2. std::optional を使う（エラー理由が不明）
// 3. 戻り値でエラーコードを返し、ポインタで値を返す（古いスタイル）
```

---

### 3. `std::optional` モナド操作 — `transform` / `and_then` / `or_else`

値が存在する場合のみ後続の処理を行うパイプラインを流れるように構築でき、ネストした `if` 文を排除できます。

```cpp
auto result = maybe_half(8)
    .transform([](int n) { return n * 3; })        // 値があれば変換
    .and_then([](int n) -> std::optional<int> {     // 値があれば処理
        return n > 0 ? std::optional{n} : std::nullopt;
    })
    .or_else([] { return std::optional{0}; });      // 値がなければ代替値
```

---

### 4. `if consteval` — コンパイル時か実行時かで分岐

コンパイル時か実行時かを直接判定でき、同一のインターフェースで実行環境に最適化された実装を切り替えられます。

```cpp
constexpr auto sq = [](int n) {
    if consteval {
        return n * n;     // コンパイル時: 二乗
    } else {
        return n + n;     // 実行時: 倍算
    }
};

static_assert(sq(5) == 25);   // 5*5 = 25
printf("sq(5) = %d\n", sq(5)); // 5+5 = 10
```

---

### 5. Deducing `this` — 明示的オブジェクトパラメータ

メンバ関数の最初の引数に `this` を明示でき、const 修飾や値/参照ごとのオーバーロードを 1 つの関数に集約してコードの重複を排除できます。

```cpp
struct Doubler {
    // `this const Doubler&` が明示的オブジェクトパラメータ
    int apply(this const Doubler& self, int n) {
        return n * 2;
    }
};

printf("result: %d\n", Doubler{}.apply(21));  // 42
```

---

## まとめ

| 機能 | ヘッダ | 概要 |
| --- | --- | --- |
| 多次元 `operator[]` | （言語機能） | `m[r, c]` で直接アクセス |
| `std::expected` | `<expected>` | 値またはエラーを返す型 |
| optional モナド操作 | `<optional>` | `transform`/`and_then`/`or_else` |
| `if consteval` | （言語機能） | コンパイル時/実行時の分岐 |
| Deducing `this` | （言語機能） | 明示的オブジェクトパラメータ |
