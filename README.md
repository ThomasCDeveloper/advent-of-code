# Advent of code

Every year I try a new language !

## Project structure

I try to use the same project structure with a somewhat same Makefile every year.

```
.
├── Makefile
├── README.md
├── .gitignore
└── solutions
    ├── 01
    │   ├── solution.lingo
    │   ├── input.txt
    │   └── test.txt
    ├── 02
    ...
```

## How to use

To use my solution you can `cd` into a specific year and run:

```bash
make run d=5 # will run day 5
```

Other make function are available:

```bash
make new # creates a new day (only useful during dev)
make test d=5 # will run day 5 with test input
make time d=5 # will run a performance test on day 5 (needs hyperfine)
```

Note that `d=5` is to specify a day. It can be ommited if you only want to run the last day.
