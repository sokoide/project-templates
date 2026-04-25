#include "sub.h"

#include <gtest/gtest.h>

TEST(Matrix2D, Subscript) {
    Matrix2D m;
    m[0, 0] = 1;
    m[1, 2] = 7;
    m[2, 1] = 5;
    EXPECT_EQ((m[0, 0]), 1);
    EXPECT_EQ((m[1, 2]), 7);
    EXPECT_EQ((m[2, 1]), 5);
}

TEST(Expected, SafeDivide) {
    auto ok = safe_divide(10, 3);
    EXPECT_TRUE(ok.has_value());
    EXPECT_EQ(ok.value(), 3);

    auto err = safe_divide(10, 0);
    EXPECT_FALSE(err.has_value());
    EXPECT_EQ(err.error(), "division by zero");
}

TEST(Optional, MonadicOps) {
    auto some = maybe_half(8).transform([](int n) { return n * 3; });
    EXPECT_TRUE(some.has_value());
    EXPECT_EQ(*some, 12);

    auto none = maybe_half(7).transform([](int n) { return n * 3; });
    EXPECT_FALSE(none.has_value());

    auto fallback = maybe_half(7).or_else([] { return std::optional{99}; });
    EXPECT_TRUE(fallback.has_value());
    EXPECT_EQ(*fallback, 99);
}

TEST(Basic, AddSub) {
    EXPECT_EQ(add(1, 2), 3);
    EXPECT_EQ(sub(5, 3), 2);
}
