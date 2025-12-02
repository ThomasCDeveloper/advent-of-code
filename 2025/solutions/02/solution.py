import sys


def parse_input(file_path):
    with open(file_path) as f:
        return [line.strip() for line in f.readlines()]


def check_num(num):
    string = str(num)
    if len(string) % 2 == 1:
        return False
    if string[: int(len(string) / 2)] == string[int(len(string) / 2) :]:
        return True
    return False


def part1(data):
    count = 0
    ranges = [[int(a.split("-")[0]), int(a.split("-")[1])] for a in data[0].split(",")]

    for [low, high] in ranges:
        for i in range(low, high + 1):
            if check_num(i):
                count += i

    return count


def check_num_length(num, i):
    token = num[:i]
    j = 0
    while j < len(num):
        if num[j : j + i] != token:
            return False
        j += len(token)
    return True


def check_num2(num):
    string = str(num)
    for i in range(1, len(string)):
        if check_num_length(string, i):
            return True
    return False


def part2(data):
    count = 0
    ranges = [[int(a.split("-")[0]), int(a.split("-")[1])] for a in data[0].split(",")]

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
    print("Part 1:", part1(data))
    print("Part 2:", part2(data))
