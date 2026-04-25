#pragma once

#include <expected>
#include <optional>
#include <string>

// C++23: multidimensional subscript operator
class Matrix2D {
    int data_[3][3]{};

  public:
    int& operator[](size_t r, size_t c) {
        return data_[r][c];
    }
    const int& operator[](size_t r, size_t c) const {
        return data_[r][c];
    }
};

// C++23: std::expected<T, E>
std::expected<int, std::string> safe_divide(int a, int b);

// C++23: std::optional monadic operations (transform, and_then, or_else)
std::optional<int> maybe_half(int n);

int add(int a, int b);
int sub(int a, int b);
