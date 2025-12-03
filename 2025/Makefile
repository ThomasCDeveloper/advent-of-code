.PHONY: run test new time

SOLUTIONS := $(shell find solutions -maxdepth 1 -type d -printf "%f\n" | grep -E '^[0-9]{2}$$' | sort)
LAST_DAY := $(shell printf "%s\n" $(SOLUTIONS) | tail -n 1)
DAY := $(shell printf "%02d" $(d))
NEXT := $(shell find solutions -maxdepth 1 -type d -printf "%f\n" | grep -E '^[0-9]{2}$$' | sort) + 1

run:
	@echo 'Running day $(if $(d),$(DAY),$(LAST_DAY))'
	@python3 -m solutions.$(if $(d),$(DAY),$(LAST_DAY)).solution

test:
	@echo 'Testing day $(if $(d),$(DAY),$(LAST_DAY))'
	@python3 -m solutions.$(if $(d),$(DAY),$(LAST_DAY)).solution test

new:
	@next=$$(printf "%02d" $$(echo $$(ls solutions | grep -E '^[0-9]{2}$$' | sort | tail -n1) + 1 | bc)); \
	mkdir -p solutions/$$next; \
	cp template/solution.py solutions/$$next/solution.py; \
	touch solutions/$$next/input.txt; \
	touch solutions/$$next/test.txt; \
	echo $$next

time:
	@echo 'Testing day $(if $(d),$(DAY),$(LAST_DAY))'
	@hyperfine --warmup 3 'python3 -m solutions.$(if $(d),$(DAY),$(LAST_DAY)).solution'