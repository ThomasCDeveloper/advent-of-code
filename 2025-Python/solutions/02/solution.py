import sys
import utils
from pathlib import Path


def check_num(num):
    s = str(num)
    n = len(s)
    if n % 2 == 1:
        return False
    return s[: n // 2] == s[n // 2 :]


def part1(ranges):
    total = 0
    for low, high in ranges:
        for i in range(low, high + 1):
            s = str(i)
            n = len(s)
            if not (n & 1) and s[: n // 2] == s[n // 2 :]:
                total += i
    return total


def check_num2(num):
    s = str(num)
    n = len(s)
    for i in range(1, n // 2 + 1):
        if n % i == 0 and s == s[:i] * (n // i):
            return True
    return False


def part2(ranges):
    total = 0
    for low, high in ranges:
        for i in range(low, high + 1):
            s = str(i)
            n = len(s)
            for k in range(1, n // 2 + 1):
                if n % k == 0 and s == s[:k] * (n // k):
                    total += i
                    break
    return total


if __name__ == "__main__":
    BASE = Path(__file__).resolve().parent
    file_path = BASE / utils.get_mode(sys.argv)
    data = [
        list(map(int, x.split("-"))) for x in utils.parse_input(file_path).split(",")
    ]
    print("Part 1:", part1(data))
    print("Part 2:", part2(data))
