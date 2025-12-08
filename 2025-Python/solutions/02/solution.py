import sys
import utils
from pathlib import Path


def check_num(num):
    s = str(num)
    n = len(s)
    if n % 2 == 1:
        return False
    return s[: n // 2] == s[n // 2 :]


def part1(data):
    ranges = [list(map(int, x.split("-"))) for x in data.split(",")]
    total = 0
    for low, high in ranges:
        for i in range(low, high + 1):
            s = str(i)
            n = len(s)
            if not (n & 1) and s[: n // 2] == s[n // 2 :]:
                total += i
    return total


def part2(data):
    ranges = [tuple(map(int, x.split("-"))) for x in data.split(",")]

    global_low = min(low for low, _ in ranges)
    global_high = max(high for _, high in ranges)
    max_digits = len(str(global_high))

    repeating_numbers = set()

    for n in range(2, max_digits + 1):
        for k in range(1, n // 2 + 1):
            if n % k != 0:
                continue

            repeat_count = n // k
            if repeat_count < 2:
                continue

            start_block = 10 ** (k - 1)
            end_block = 10**k

            for block in range(start_block, end_block):
                s = str(block) * repeat_count
                if len(s) != n:
                    continue

                value = int(s)
                if value > global_high:
                    break
                if value >= global_low:
                    repeating_numbers.add(value)

    total = 0
    for value in repeating_numbers:
        for low, high in ranges:
            if low <= value <= high:
                total += value
                break

    return total


if __name__ == "__main__":
    BASE = Path(__file__).resolve().parent
    file_path = BASE / utils.get_mode(sys.argv)
    data = utils.parse_input(file_path)
    print("Part 1:", part1(data))
    print("Part 2:", part2(data))
