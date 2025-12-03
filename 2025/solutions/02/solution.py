import sys
import utils
from pathlib import Path


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
    BASE = Path(__file__).resolve().parent
    file_path = BASE / utils.get_mode(sys.argv)
    data = utils.parse_input_lines(file_path)
    print("Part 1:", part1(data))
    print("Part 2:", part2(data))
