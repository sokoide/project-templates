#pragma once

#include <compare>
#include <concepts>
#include <span>
#include <string>
#include <vector>

// C++20: concepts
template <std::integral T> T generic_add(T a, T b);
template <std::floating_point T> T generic_add(T a, T b);

// C++20: three-way comparison (spaceship operator)
struct Point {
    int x, y;
    auto operator<=>(const Point &) const = default;
};

// C++20: std::span
int span_sum(std::span<const int> s);

// C++20: designated initializers
struct Config {
    int width = 0;
    int height = 0;
    bool fullscreen = false;
};
Config make_default_config();

// C++20: consteval
consteval int square(int n) { return n * n; }

// C++20: ranges
std::vector<int> filter_even(const std::vector<int> &input);
std::vector<int> transform_double(const std::vector<int> &input);
