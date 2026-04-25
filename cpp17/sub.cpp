#include "sub.h"
#include <algorithm>

std::optional<int> find_index(const std::vector<int> &v, int target) {
    auto it = std::find(v.begin(), v.end(), target);
    if (it != v.end())
        return static_cast<int>(it - v.begin());
    return std::nullopt;
}

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

DivResult int_div(int a, int b) { return {a / b, a % b}; }
