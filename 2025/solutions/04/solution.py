import sys
import utils
from pathlib import Path


def part1(data: list[str]) -> int:
    count = 0
    rolls = {(x, y) for y, l in enumerate(data) for x, c in enumerate(l) if c == "@"}
    for x, y in rolls:
        sum = -1
        for dx in [-1, 0, 1]:
            for dy in [-1, 0, 1]:
                if (x + dx, y + dy) in rolls:
                    sum += 1
        if sum < 4:
            count += 1
    return count


def part2(data: list[str]) -> int:
    rolls = {(x, y) for y, l in enumerate(data) for x, c in enumerate(l) if c == "@"}
    count = 0
    changed = 1
    while changed:
        to_remove = set()
        changed = False
        for x, y in rolls:
            sum = -1
            for dx in [-1, 0, 1]:
                for dy in [-1, 0, 1]:
                    if (x + dx, y + dy) in rolls:
                        sum += 1
            if sum < 4:
                to_remove.add((x, y))
        if to_remove:
            rolls -= to_remove
            count += len(to_remove)
            changed = True

    return count


if __name__ == "__main__":
    BASE = Path(__file__).resolve().parent
    file_path = BASE / utils.get_mode(sys.argv)
    data = utils.parse_input_lines(file_path)
    print("Part 1:", part1(data))
    print("Part 2:", part2(data))
