import sys


def parse_input(file_path):
    with open(file_path) as f:
        return [line.strip() for line in f.readlines()]


def check_num(num):
    string = str(num)
    if len(string) % 2 == 1:
        return False
    if string == string[int(len(string) / 2) :] * 2:
        return True
    return False


def part1(ranges):
    count = 0
    for [low, high] in ranges:
        for i in range(low, high + 1):
            if check_num(i):
                count += i

    return count


def check_num_length(num, i):
    token = num[:i]
    if num == token * (len(num) // i):
        return True
    return False


def check_num2(num):
    string = str(num)
    for i in range(1, len(string)):
        if check_num_length(string, i):
            return True
    return False


def part2(ranges):
    count = 0
    for [low, high] in ranges:
        for i in range(low, high + 1):
            if check_num2(i):
                count += i

    return count


if __name__ == "__main__":
    if len(sys.argv) != 2 or sys.argv[1] not in ("input", "test"):
        print("Usage: python solution.py [input|test]")
        sys.exit(1)

    file_path = (
        "./solutions/02/input.txt"
        if sys.argv[1] == "input"
        else "./solutions/02/test.txt"
    )

    data = parse_input(file_path)
    ranges = [[int(a.split("-")[0]), int(a.split("-")[1])] for a in data[0].split(",")]
    print("Part 1:", part1(ranges))
    print("Part 2:", part2(ranges))
