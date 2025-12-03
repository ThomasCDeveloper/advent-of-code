import sys


def parse_input(file_path):
    with open(file_path) as f:
        return [line.strip() for line in f.readlines()]


def part1(data):
    return 0  # TODO


def part2(data):
    return 0  # TODO


if __name__ == "__main__":
    file_path = "input.txt"
    if sys.argv[1] == "test":
        file_path = "test.txt"

    data = parse_input(file_path)
    print("Part 1:", part1(data))
    print("Part 2:", part2(data))
