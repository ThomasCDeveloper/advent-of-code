import sys
import utils
from pathlib import Path


def get_cell_status(x, y, data):
    if y < 0 or y > len(data) - 1 or x < 0 or x > len(data[0]) - 1:
        return "."
    else:
        return data[y][x]


def part1(data: list[str]) -> int:
    count = 0
    for y in range(len(data)):
        for x in range(len(data[0])):
            if data[y][x] == "@":
                sum = -1
                for dx in [-1, 0, 1]:
                    for dy in [-1, 0, 1]:
                        if get_cell_status(x + dx, y + dy, data) == "@":
                            sum += 1
                if sum < 4:
                    count += 1
    return count


def part2(data: list[str]) -> int:
    count = 0
    changed = True
    while changed:
        changed = False
        for y in range(len(data)):
            for x in range(len(data[0])):
                if data[y][x] == "@":
                    sum = -1
                    for dx in [-1, 0, 1]:
                        for dy in [-1, 0, 1]:
                            if get_cell_status(x + dx, y + dy, data) == "@":
                                sum += 1
                    if sum < 4:
                        changed = True
                        data[y] = data[y][:x] + "." + data[y][x + 1 :]
                        count += 1
    return count


if __name__ == "__main__":
    BASE = Path(__file__).resolve().parent
    file_path = BASE / utils.get_mode(sys.argv)
    data = utils.parse_input_lines(file_path)
    print("Part 1:", part1(data))
    print("Part 2:", part2(data))
