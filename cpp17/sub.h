#pragma once

#include <optional>
#include <string>
#include <string_view>
#include <type_traits>
#include <variant>
#include <vector>

// C++17: std::optional
std::optional<int> find_index(const std::vector<int>& v, int target);

// C++17: std::variant (type-safe union)
using Value = std::variant<int, double, std::string>;
std::string value_type_name(const Value& v);

// C++17: std::string_view
size_t count_words(std::string_view sv);

// C++17: structured bindings helper
struct DivResult {
    int quotient;
    int remainder;
};
DivResult int_div(int a, int b);

// C++17: fold expressions
template <typename... Args>
auto fold_sum(Args... args) {
    return (args + ...);
}

// C++17: if constexpr
template <typename T>
std::string stringify(T val) {
    if constexpr (std::is_integral_v<T>) {
        return "int:" + std::to_string(val);
    } else {
        return "float:" + std::to_string(val);
    }
}
