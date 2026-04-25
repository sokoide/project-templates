SUBDIRS := $(sort $(dir $(wildcard */Makefile)))

.PHONY: all clean test run runtest $(SUBDIRS)

all: $(SUBDIRS)

$(SUBDIRS):
	$(MAKE) -C $@ $(MAKECMDGOALS)

clean: $(SUBDIRS)
test: $(SUBDIRS)
run: $(SUBDIRS)
runtest: $(SUBDIRS)
