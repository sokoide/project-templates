#include "sub.h"

int add(int a, int b) {
    return a + b;
}
int sub(int a, int b) {
    return a - b;
}

std::expected<int, std::string> safe_divide(int a, int b) {
    if (b == 0)
        return std::unexpected("division by zero");
    return a / b;
}

std::optional<int> maybe_half(int n) {
    if (n % 2 != 0)
        return std::nullopt;
    return n / 2;
}
