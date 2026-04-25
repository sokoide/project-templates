#pragma once

#include <array>
#include <functional>
#include <memory>
#include <string>
#include <vector>

// C++11: enum class
enum class Color { Red, Green, Blue };
std::string color_name(Color c);

// C++11: constexpr
constexpr int factorial(int n) {
    return n <= 1 ? 1 : n * factorial(n - 1);
}

// C++11: std::array
int array_sum(const std::array<int, 5>& arr);

// C++11: smart pointers
std::unique_ptr<int> make_unique_val(int v);
std::shared_ptr<std::string> make_shared_str(const std::string& s);

// C++11: std::function + lambda
std::function<int(int)> make_adder(int base);

// C++11: move semantics
class Buffer {
    int* data_;
    size_t size_;

  public:
    Buffer(size_t sz);
    ~Buffer();
    Buffer(Buffer&& other) noexcept;
    Buffer& operator=(Buffer&& other) noexcept;
    Buffer(const Buffer&) = delete;
    Buffer& operator=(const Buffer&) = delete;
    size_t size() const;
};

// C++11: initializer_list
std::vector<int> make_sorted(std::initializer_list<int> vals);
