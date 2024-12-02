.PHONY: all
all: build

# Build the project
.PHONY: build
build:
	dune build

.PHONY: watch
watch:
	DUNE_CONFIG__GLOBAL_LOCK=disabled dune build @all --watch

# Clean build artifacts
.PHONY: clean
clean:
	dune clean

# Run the project for a specific day
.PHONY: run
run:
	@if [ -z "$(DAY)" ]; then \
		read -p "Enter the day number to run: " DAY; \
	fi; \
	_build/default/bin/main.exe $$DAY

# Run and time the project for a specific day
.PHONY: time
time:
	@if [ -z "$(DAY)" ]; then \
		read -p "Enter the day number to run: " DAY; \
	fi; \
	hyperfine "_build/default/bin/main.exe $$DAY"