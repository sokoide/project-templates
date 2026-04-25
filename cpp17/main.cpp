#include "sub.h"
#include <cstdio>
#include <map>
#include <optional>
#include <tuple>

int main() {
    //--- C++17: std::optional ---
    std::vector<int> v = {10, 20, 30, 40};
    auto idx = find_index(v, 30);
    printf("find 30: ");
    if (idx)
        printf("index=%d\n", *idx);
    else
        printf("not found\n");
    auto miss = find_index(v, 99);
    printf("find 99: ");
    if (miss)
        printf("index=%d\n", *miss);
    else
        printf("not found\n");

    //--- C++17: std::variant ---
    Value vi = 42;
    Value vd = 3.14;
    Value vs = std::string("hello");
    printf("variant types: %s %s %s\n", value_type_name(vi).c_str(),
           value_type_name(vd).c_str(), value_type_name(vs).c_str());

    //--- C++17: std::string_view ---
    printf("word count: %zu\n", count_words("hello world foo bar"));

    //--- C++17: structured bindings (struct) ---
    auto [q, r] = int_div(17, 5);
    printf("17/5 = %d remainder %d\n", q, r);

    //--- C++17: structured bindings (map) ---
    std::map<std::string, int> scores;
    scores["alice"] = 90;
    scores["bob"] = 85;
    for (auto &[name, score] : scores) {
        printf("  %s: %d\n", name.c_str(), score);
    }

    //--- C++17: structured bindings (tuple) ---
    auto [x, y, z] = std::make_tuple(1, 2.5, "three");
    printf("tuple: %d %.1f %s\n", x, y, z);

    //--- C++17: fold expressions ---
    printf("fold_sum(1..5)=%d\n", fold_sum(1, 2, 3, 4, 5));
    printf("fold_sum(1.1,2.2,3.3)=%.1f\n", fold_sum(1.1, 2.2, 3.3));

    //--- C++17: if constexpr ---
    printf("stringify(42)=%s\n", stringify(42).c_str());
    printf("stringify(3.14)=%s\n", stringify(3.14).c_str());

    //--- C++17: CTAD (Class Template Argument Deduction) ---
    std::pair p{1, 2};
    std::optional opt{99};
    printf("CTAD pair: (%d,%d)\n", p.first, p.second);
    printf("CTAD optional: %d\n", *opt);

    return 0;
}
