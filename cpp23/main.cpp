#include "sub.h"
#include <cstdio>
#include <utility>

int main() {
    //--- C++23: multidimensional operator[] ---
    Matrix2D m;
    for (int i = 0; i < 3; i++)
        for (int j = 0; j < 3; j++)
            m[i, j] = i * 3 + j + 1;
    printf("m[0,0]=%d m[1,2]=%d m[2,2]=%d\n", m[0, 0], m[1, 2], m[2, 2]);

    //--- C++23: std::expected ---
    auto r1 = safe_divide(10, 3);
    auto r2 = safe_divide(10, 0);
    if (r1)
        printf("10/3=%d\n", *r1);
    if (!r2)
        printf("10/0 error: %s\n", r2.error().c_str());

    //--- C++23: std::optional monadic ops ---
    auto ok = maybe_half(8)
                  .transform([](int n) { return n * 3; })
                  .and_then([](int n) -> std::optional<int> {
                      return n > 0 ? std::optional{n} : std::nullopt;
                  })
                  .or_else([] { return std::optional{0}; });
    printf("maybe_half(8)->transform(*3)->and_then->or_else: %d\n", *ok);

    auto ng = maybe_half(7)
                  .transform([](int n) { return n * 3; })
                  .or_else([] { return std::optional{-1}; });
    printf("maybe_half(7)->transform->or_else(-1): %d\n", *ng);

    //--- C++23: if consteval ---
    constexpr auto sq = [](int n) {
        if consteval {
            return n * n;
        } else {
            return n + n;
        }
    };
    printf("sq(5) runtime: %d\n", sq(5));
    static_assert(sq(5) == 25);

    //--- C++23: deducing this ---
    struct Doubler {
        int apply(this const Doubler &, int n) { return n * 2; }
    };
    printf("deducing this: %d\n", Doubler{}.apply(21));

    return 0;
}
