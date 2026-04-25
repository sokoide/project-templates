OUTDIR = ./build
TARGET = cpp20.out
TEST_TARGET = cpp20_test.out
PCH = pch.hpp

GTEST_DIR = /opt/homebrew/opt/googletest

HEADERS := $(shell find . -type f -name '*.h')
SRCS = main.cpp \
	sub.cpp
TEST_SRCS = test_main.cpp
TEST_OBJS_SRCS = sub.cpp

INCLUDEPCH = $(OUTDIR)/$(PCH)
BUILTPCH = $(OUTDIR)/pch.hpp.pch
OBJS = $(SRCS:.cpp=.o)
DEPS = $(SRCS:.cpp=.d) $(TEST_SRCS:.cpp=.d) $(TEST_OBJS_SRCS:.cpp=.d)
OUTOBJS = $(addprefix $(OUTDIR)/,$(OBJS))
OUTDEPS = $(addprefix $(OUTDIR)/,$(DEPS))
TEST_OUTOBJS = $(addprefix $(OUTDIR)/,$(TEST_SRCS:.cpp=.o) $(TEST_OBJS_SRCS:.cpp=.o))

CXX = clang++
CXXFLAGS = -std=c++20 -Wall -g -I. -I/usr/local/include -O0
TEST_CXXFLAGS = $(CXXFLAGS) -I$(GTEST_DIR)/include
TEST_LDFLAGS = -L$(GTEST_DIR)/lib -lgtest -lgtest_main

default: all

all: $(TARGET)

copy:
	cat main.cpp|pbcopy

lldb: $(TARGET)
	lldb $(OUTDIR)/$(TARGET)

$(TARGET): $(BUILTPCH) link

link: $(OUTOBJS)
	$(CXX) $(CXXFLAGS) $^ -o $(OUTDIR)/$(TARGET)

test: $(TEST_TARGET)

$(TEST_TARGET): $(BUILTPCH) $(TEST_OUTOBJS)
	$(CXX) $(TEST_OUTOBJS) -o $(OUTDIR)/$(TEST_TARGET) $(TEST_LDFLAGS)

runtest: $(TEST_TARGET)
	$(OUTDIR)/$(TEST_TARGET)

diag: $(TARGET)
	objdump -p $(OUTDIR)/$(TARGET)

run: $(TARGET)
	$(OUTDIR)/$(TARGET)

clean:
	rm -rf $(BUILTPCH) $(OUTOBJS) $(TEST_OUTOBJS) $(OUTDEPS) $(OUTDIR)/$(TARGET)*

$(OUTDIR):
	mkdir -p $(OUTDIR)

$(OUTDIR)/test_%.o: test_%.cpp | $(BUILTPCH) $(OUTDIR)/test_%.d
	$(CXX) $(TEST_CXXFLAGS) -include $(INCLUDEPCH) -c $< -o $@

$(OUTDIR)/%.o: %.cpp | $(BUILTPCH) $(OUTDIR)/%.d
	$(CXX) $(CXXFLAGS) -include $(INCLUDEPCH) -c $< -o $@

$(BUILTPCH): $(PCH) | $(OUTDIR)
	$(CXX) $(CXXFLAGS) -x c++-header $(PCH) -o $@

$(OUTDIR)/test_%.d: test_%.cpp $(HEADERS) | $(OUTDIR)
	$(CXX) $(TEST_CXXFLAGS) -MM -MP $< -MF $@

$(OUTDIR)/%.d: %.cpp $(HEADERS) | $(OUTDIR)
	$(CXX) $(CXXFLAGS) -MM -MP $< -MF $@
