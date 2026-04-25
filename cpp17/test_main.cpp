#include "sub.h"

#include <gtest/gtest.h>
#include <map>
#include <tuple>

TEST(Optional, FindIndex) {
    std::vector<int> v = {10, 20, 30};
    auto found = find_index(v, 20);
    ASSERT_TRUE(found.has_value());
    EXPECT_EQ(*found, 1);

    auto miss = find_index(v, 99);
    EXPECT_FALSE(miss.has_value());
}

TEST(Variant, TypeName) {
    EXPECT_EQ(value_type_name(Value(42)), "int");
    EXPECT_EQ(value_type_name(Value(3.14)), "double");
    EXPECT_EQ(value_type_name(Value(std::string("hi"))), "string");
}

TEST(StringView, CountWords) {
    EXPECT_EQ(count_words("hello world"), 2u);
    EXPECT_EQ(count_words("  one  two three  "), 3u);
    EXPECT_EQ(count_words(""), 0u);
}

TEST(StructuredBindings, DivResult) {
    auto [q, r] = int_div(17, 5);
    EXPECT_EQ(q, 3);
    EXPECT_EQ(r, 2);
}

TEST(StructuredBindings, MapIter) {
    std::map<int, std::string> m;
    m[1] = "a";
    m[2] = "b";
    int sum = 0;
    for (auto& [k, v] : m) sum += k;
    EXPECT_EQ(sum, 3);
}

TEST(StructuredBindings, Tuple) {
    auto [a, b, c] = std::make_tuple(1, 2.5, std::string("hi"));
    EXPECT_EQ(a, 1);
    EXPECT_DOUBLE_EQ(b, 2.5);
    EXPECT_EQ(c, "hi");
}

TEST(FoldExpr, Sum) {
    EXPECT_EQ(fold_sum(1, 2, 3), 6);
    EXPECT_EQ(fold_sum(10, 20, 30, 40), 100);
}

TEST(IfConstexpr, Stringify) {
    EXPECT_EQ(stringify(42), "int:42");
    EXPECT_EQ(stringify(3.14), "float:3.140000");
}
