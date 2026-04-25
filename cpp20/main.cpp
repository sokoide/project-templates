#include "sub.h"
#include <cstdio>
#include <ranges>
#include <vector>

int main() {
    //--- C++20: concepts ---
    printf("add<int>(1,2)=%d\n", generic_add(1, 2));
    printf("add<double>(1.5,2.5)=%.1f\n", generic_add(1.5, 2.5));

    //--- C++20: three-way comparison (<=>) ---
    Point a{1, 2}, b{3, 4}, c{1, 2};
    auto cmp = (a <=> b);
    printf("a<=>b: %s\n", std::is_lt(cmp) ? "less" : std::is_gt(cmp) ? "greater" : "equal");
    printf("a==c: %d\n", a == c);

    //--- C++20: std::span ---
    int arr[] = {10, 20, 30, 40, 50};
    printf("span_sum(arr)=%d\n", span_sum(arr));
    std::vector<int> v = {1, 2, 3};
    printf("span_sum(vector)=%d\n", span_sum(v));

    //--- C++20: designated initializers ---
    Config cfg = make_default_config();
    Config custom{.width = 1920, .height = 1080, .fullscreen = true};
    printf("default config: %dx%d fs=%d\n", cfg.width, cfg.height, cfg.fullscreen);
    printf("custom config: %dx%d fs=%d\n", custom.width, custom.height, custom.fullscreen);

    //--- C++20: consteval ---
    static_assert(square(5) == 25);
    printf("square(5)=%d\n", square(5));

    //--- C++20: ranges algorithms ---
    std::vector<int> nums = {1, 2, 3, 4, 5, 6};
    auto evens = filter_even(nums);
    printf("evens:");
    for (auto n : evens) printf(" %d", n);
    printf("\n");

    auto doubled = transform_double(nums);
    printf("doubled:");
    for (auto n : doubled) printf(" %d", n);
    printf("\n");

    //--- C++20: ranges views pipeline ---
    namespace rv = std::views;
    auto result = nums | rv::filter([](int n) { return n % 2 == 0; })
                       | rv::transform([](int n) { return n * n; });
    printf("squared evens:");
    for (auto n : result) printf(" %d", n);
    printf("\n");

    //--- C++20: ranges views::take & iota ---
    auto first3 = rv::iota(10, 20) | rv::take(3);
    printf("iota(10,20)|take(3):");
    for (auto n : first3) printf(" %d", n);
    printf("\n");

    return 0;
}
