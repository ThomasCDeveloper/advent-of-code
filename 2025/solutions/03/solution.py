import sys
import utils
from pathlib import Path


def parse_input(file_path):
    with open(file_path) as f:
        return [line.strip() for line in f.readlines()]


def get_highest_digit_and_remainder(digit_line, offset):
    digits = [int(d) for d in digit_line]
    max = 0
    for i, d in enumerate(digits):
        if i == len(digits) - offset:
            break
        if d > max:
            max = d
    for i in range(len(digit_line)):
        if digit_line[i] == str(max):
            return [max, digit_line[i + 1 :]]


def part1(data: list[str]) -> int:
    total = 0
    for line in data:
        first, remainder = get_highest_digit_and_remainder(line, 1)
        second, remainder = get_highest_digit_and_remainder(remainder, 0)
        total += 10 * first + second

    return total


def part2(data: list[str]) -> int:
    total = 0
    for line in data:
        subtotal = 0
        remaining = line
        for i in range(12):
            digit, remaining = get_highest_digit_and_remainder(remaining, 11 - i)
            subtotal = subtotal * 10 + digit
        total += subtotal
    return total


if __name__ == "__main__":
    BASE = Path(__file__).resolve().parent
    file_path = BASE / utils.get_mode(sys.argv)
    data = utils.parse_input_lines(file_path)
    print("Part 1:", part1(data))
    print("Part 2:", part2(data))
