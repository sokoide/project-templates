#include "sub.h"

#include <gtest/gtest.h>

TEST(EnumClass, ColorName) {
    EXPECT_EQ(color_name(Color::Red), "Red");
    EXPECT_EQ(color_name(Color::Green), "Green");
    EXPECT_EQ(color_name(Color::Blue), "Blue");
}

TEST(Constexpr, Factorial) {
    EXPECT_EQ(factorial(0), 1);
    EXPECT_EQ(factorial(1), 1);
    EXPECT_EQ(factorial(5), 120);
    EXPECT_EQ(factorial(10), 3628800);
}

TEST(Array, Sum) {
    std::array<int, 5> arr = {{1, 2, 3, 4, 5}};
    EXPECT_EQ(array_sum(arr), 15);
}

TEST(SmartPtr, UniquePtr) {
    auto p = make_unique_val(42);
    ASSERT_TRUE(p != nullptr);
    EXPECT_EQ(*p, 42);
}

TEST(SmartPtr, SharedPtr) {
    auto p = make_shared_str("hello");
    EXPECT_EQ(*p, "hello");
    EXPECT_EQ(p.use_count(), 1);
    auto p2 = p;
    EXPECT_EQ(p.use_count(), 2);
}

TEST(Lambda, Adder) {
    auto add10 = make_adder(10);
    EXPECT_EQ(add10(5), 15);
    EXPECT_EQ(add10(0), 10);
}

TEST(MoveSemantics, Buffer) {
    Buffer a(50);
    EXPECT_EQ(a.size(), 50u);
    Buffer b = std::move(a);
    EXPECT_EQ(b.size(), 50u);
    EXPECT_EQ(a.size(), 0u);
}

TEST(InitializerList, Sorted) {
    auto v = make_sorted({5, 3, 1, 4, 2});
    ASSERT_EQ(v.size(), 5u);
    EXPECT_EQ(v[0], 1);
    EXPECT_EQ(v[4], 5);
}
