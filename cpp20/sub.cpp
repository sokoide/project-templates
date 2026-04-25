#include "sub.h"
#include <algorithm>
#include <ranges>

template <std::integral T> T generic_add(T a, T b) { return a + b; }
template <std::floating_point T> T generic_add(T a, T b) { return a + b; }

int span_sum(std::span<const int> s) {
    int sum = 0;
    for (auto v : s)
        sum += v;
    return sum;
}

Config make_default_config() {
    return {.width = 800, .height = 600, .fullscreen = false};
}

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

// explicit template instantiations for linking
template int generic_add<int>(int, int);
template double generic_add<double>(double, double);
