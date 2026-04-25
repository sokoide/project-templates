#include "sub.h"
#include <gtest/gtest.h>
#include <vector>

TEST(Concepts, GenericAdd) {
    EXPECT_EQ(generic_add(1, 2), 3);
    EXPECT_EQ(generic_add(-1, 1), 0);
    EXPECT_DOUBLE_EQ(generic_add(1.5, 2.5), 4.0);
}

TEST(Spaceship, PointComparison) {
    Point a{1, 2}, b{3, 4}, c{1, 2};
    EXPECT_TRUE(a == c);
    EXPECT_TRUE(a != b);
    EXPECT_TRUE(a < b);
    EXPECT_FALSE(b < a);
}

TEST(Span, Sum) {
    int arr[] = {10, 20, 30};
    EXPECT_EQ(span_sum(arr), 60);

    std::vector<int> v = {1, 2, 3, 4};
    EXPECT_EQ(span_sum(v), 10);
}

TEST(DesignatedInit, Config) {
    Config cfg = make_default_config();
    EXPECT_EQ(cfg.width, 800);
    EXPECT_EQ(cfg.height, 600);
    EXPECT_FALSE(cfg.fullscreen);

    Config custom{.width = 1920, .height = 1080, .fullscreen = true};
    EXPECT_EQ(custom.width, 1920);
    EXPECT_TRUE(custom.fullscreen);
}

TEST(Ranges, FilterEven) {
    std::vector<int> input = {1, 2, 3, 4, 5, 6};
    auto result = filter_even(input);
    ASSERT_EQ(result.size(), 3u);
    EXPECT_EQ(result[0], 2);
    EXPECT_EQ(result[1], 4);
    EXPECT_EQ(result[2], 6);
}

TEST(Ranges, TransformDouble) {
    std::vector<int> input = {1, 2, 3};
    auto result = transform_double(input);
    ASSERT_EQ(result.size(), 3u);
    EXPECT_EQ(result[0], 2);
    EXPECT_EQ(result[1], 4);
    EXPECT_EQ(result[2], 6);
}
