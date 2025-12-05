import sys
import utils
from pathlib import Path


def part1(data):
    ranges, items = data.split("\n\n")
    ranges = [list(map(int, x.split("-"))) for x in ranges.split("\n")]
    items = list(map(int, items.split("\n")))

    count = 0
    for item in items:
        validated = False
        for start, end in ranges:
            if validated:
                break
            if item >= start and item <= end:
                count += 1
                validated = True

    return count


def part2(data):
    ranges = [list(map(int, x.split("-"))) for x in data.split("\n\n")[0].split("\n")]
    ranges = sorted(ranges, key=lambda x: x[0])
    count = 0

    current = -1
    for start, end in ranges:
        if current >= start:
            start = current + 1
        if start <= end:
            count += end - start + 1
        current = max(current, end)

    return count


if __name__ == "__main__":
    BASE = Path(__file__).resolve().parent
    file_path = BASE / utils.get_mode(sys.argv)
    data = utils.parse_input(file_path)
    print("Part 1:", part1(data))
    print("Part 2:", part2(data))
