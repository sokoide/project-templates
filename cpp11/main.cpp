#include "sub.h"
#include <cstdio>
#include <map>
#include <vector>

int main() {
    //--- C++11: auto ---
    auto x = 42;
    auto name = std::string("hello");
    printf("auto: x=%d name=%s\n", x, name.c_str());

    //--- C++11: range-based for ---
    std::vector<int> v = {1, 2, 3, 4, 5};
    printf("range-for:");
    for (auto n : v)
        printf(" %d", n);
    printf("\n");

    //--- C++11: lambda ---
    auto greet = [](const char *s) { printf("lambda: %s\n", s); };
    greet("hello");
    int capture = 10;
    auto add = [capture](int n) { return n + capture; };
    printf("lambda capture: %d\n", add(5));

    //--- C++11: enum class ---
    Color c = Color::Green;
    printf("enum class: %s\n", color_name(c).c_str());

    //--- C++11: constexpr ---
    printf("constexpr factorial(5)=%d\n", factorial(5));

    //--- C++11: std::array ---
    std::array<int, 5> arr = {{10, 20, 30, 40, 50}};
    printf("array_sum: %d\n", array_sum(arr));

    //--- C++11: unique_ptr / shared_ptr ---
    auto up = make_unique_val(42);
    printf("unique_ptr: %d\n", *up);
    auto sp = make_shared_str("smart");
    printf("shared_ptr: %s use_count=%ld\n", sp->c_str(), sp.use_count());

    //--- C++11: std::function ---
    auto add5 = make_adder(5);
    printf("adder(5)(10)=%d\n", add5(10));

    //--- C++11: move semantics ---
    Buffer b1(100);
    printf("b1 size=%zu\n", b1.size());
    Buffer b2 = std::move(b1);
    printf("after move: b2 size=%zu\n", b2.size());

    //--- C++11: initializer_list ---
    auto sorted = make_sorted({5, 3, 1, 4, 2});
    printf("sorted:");
    for (auto n : sorted)
        printf(" %d", n);
    printf("\n");

    //--- C++11: range-based for with map ---
    std::map<std::string, int> m;
    m["a"] = 1;
    m["b"] = 2;
    printf("map:");
    for (auto &kv : m)
        printf(" %s=%d", kv.first.c_str(), kv.second);
    printf("\n");

    return 0;
}
