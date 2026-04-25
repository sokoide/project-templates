#include "sub.h"
#include <algorithm>

std::string color_name(Color c) {
    switch (c) {
    case Color::Red:   return "Red";
    case Color::Green: return "Green";
    case Color::Blue:  return "Blue";
    }
    return "Unknown";
}

int array_sum(const std::array<int, 5> &arr) {
    int sum = 0;
    for (auto v : arr)
        sum += v;
    return sum;
}

std::unique_ptr<int> make_unique_val(int v) {
    return std::unique_ptr<int>(new int(v));
}

std::shared_ptr<std::string> make_shared_str(const std::string &s) {
    return std::make_shared<std::string>(s);
}

std::function<int(int)> make_adder(int base) {
    return [base](int x) { return base + x; };
}

Buffer::Buffer(size_t sz) : data_(new int[sz]), size_(sz) {}
Buffer::~Buffer() { delete[] data_; }
Buffer::Buffer(Buffer &&other) noexcept
    : data_(other.data_), size_(other.size_) {
    other.data_ = nullptr;
    other.size_ = 0;
}
Buffer &Buffer::operator=(Buffer &&other) noexcept {
    if (this != &other) {
        delete[] data_;
        data_ = other.data_;
        size_ = other.size_;
        other.data_ = nullptr;
        other.size_ = 0;
    }
    return *this;
}
size_t Buffer::size() const { return size_; }

std::vector<int> make_sorted(std::initializer_list<int> vals) {
    std::vector<int> v(vals);
    std::sort(v.begin(), v.end());
    return v;
}
