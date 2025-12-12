.PHONY: run test new time

SOLUTIONS := $(shell find solutions -maxdepth 1 -type d -printf "%f\n" | grep -E '^[0-9]{2}$$' | sort)
LAST_DAY := $(shell printf "%s\n" $(SOLUTIONS) | tail -n 1)
DAY := $(shell printf "%02d" $(d))

run:
	@echo "Running day $(if $(d),$(DAY),$(LAST_DAY))"
	@dune exec ./solutions/$(if $(d),$(DAY),$(LAST_DAY))/solution.exe input $(if $(d),$(DAY),$(LAST_DAY))

test:
	@echo "Testing day $(if $(d),$(DAY),$(LAST_DAY))"
	@dune exec ./solutions/$(if $(d),$(DAY),$(LAST_DAY))/solution.exe test $(if $(d),$(DAY),$(LAST_DAY))

new:
	@next=$$(printf "%02d" $$(echo $$(ls solutions | grep -E '^[0-9]{2}$$' | sort | tail -n1) + 1 | bc)); \
	mkdir -p solutions/$$next; \
	cp template/solution.ml solutions/$$next/solution.ml; \
	cp template/dune solutions/$$next/dune; \
	touch solutions/$$next/input.txt; \
	touch solutions/$$next/test.txt; \
	echo $$next

time:
	@echo "Timing day $(if $(d),$(DAY),$(LAST_DAY))"
	@hyperfine --warmup 3 'dune exec ./solutions/$(if $(d),$(DAY),$(LAST_DAY))/solution.exe input $(if $(d),$(DAY),$(LAST_DAY))'
