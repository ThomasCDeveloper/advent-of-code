.PHONY: run test new time

SOLUTIONS := $(shell find solutions -maxdepth 1 -type d -printf "%f\n" | grep -E '^[0-9]{2}$$' | sort)
LAST_DAY := $(shell printf "%s\n" $(SOLUTIONS) | tail -n 1)
DAY := $(shell printf "%02d" $(d))

run:
	@echo 'Running day $(if $(d),$(DAY),$(LAST_DAY))'
	@go run solutions/$(if $(d),$(DAY),$(LAST_DAY))/solution.go run $(if $(d),$(DAY),$(LAST_DAY))

test:
	@echo 'Testing day $(if $(d),$(DAY),$(LAST_DAY))'
	@go run solutions/$(if $(d),$(DAY),$(LAST_DAY))/solution.go test $(if $(d),$(DAY),$(LAST_DAY))

new:
	@next=$$(printf "%02d" $$(echo $$(ls solutions | grep -E '^[0-9]{2}$$' | sort | tail -n1) + 1 | bc)); \
	mkdir -p solutions/$$next; \
	cp template/solution.go solutions/$$next/solution.go; \
	touch solutions/$$next/input.txt; \
	touch solutions/$$next/test.txt; \
	echo $$next

time:
	@echo 'Timing day $(if $(d),$(DAY),$(LAST_DAY))'
	@go build -o solutions/$(if $(d),$(DAY),$(LAST_DAY))/solution solutions/$(if $(d),$(DAY),$(LAST_DAY))/solution.go
	@hyperfine --warmup 3 './solutions/$(if $(d),$(DAY),$(LAST_DAY))/solution run $(if $(d),$(DAY),$(LAST_DAY))'
