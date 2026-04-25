SUBDIRS := $(sort $(dir $(wildcard */Makefile)))

.PHONY: all clean test run runtest $(SUBDIRS) format

all: $(SUBDIRS)

$(SUBDIRS):
	$(MAKE) -C $@ $(MAKECMDGOALS)

clean: $(SUBDIRS)
test: $(SUBDIRS)
run: $(SUBDIRS)
runtest: $(SUBDIRS)

# use pnpm if available, otherwise npx
RUNNER := $(shell command -v pnpm >/dev/null 2>&1 && echo "pnpm dlx" || echo "npx")
EXEC := $(shell command -v pnpm >/dev/null 2>&1 && echo "pnpm exec" || echo "npx")

format:
	@echo "Formatting markdown files using $(RUNNER)..."
	$(RUNNER) markdownlint-cli "**/*.md" --ignore "conductor/**" --ignore "CLAUDE.md" --ignore "node_modules/**" --ignore ".gomodcache/**" --fix
	$(EXEC) textlint --fix "*.md" "**/*.md" 
