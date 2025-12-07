import sys
import utils
import functools
from pathlib import Path


def part1(data: list[str]):
    for y, li in enumerate(data[1:]):
        for x, co in enumerate(li):
            if data[y - 1][x] in ["|", "S"]:
                if data[y][x] != "^":
                    data[y] = data[y][:x] + "|" + data[y][x + 1 :]
                else:
                    data[y] = data[y][: x - 1] + "|^|" + data[y][x + 2 :]
    total = 0
    for y, li in enumerate(data[1:]):
        for x, co in enumerate(li):
            if data[y][x] == "^":
                if data[y - 1][x] == "|":
                    total += 1
    return total


dim_x = 0
dim_y = 0
splitters = set()


@functools.lru_cache()
def get_number_of_timelines(start: int) -> int:
    y = start // dim_y

    if y > dim_y:
        return 0
    if start + dim_x in splitters:
        return (
            1
            + get_number_of_timelines(start + dim_x + 1)
            + get_number_of_timelines(start + dim_x - 1)
        )

    return get_number_of_timelines(start + dim_x)


def part2(data):
    start = 0
    for y, li in enumerate(data):
        for x, co in enumerate(li):
            if data[y][x] == "S":
                start = x + dim_x * y
            if data[y][x] == "^":
                splitters.add(x + dim_x * y)
    return 1 + get_number_of_timelines(start)


if __name__ == "__main__":
    BASE = Path(__file__).resolve().parent
    file_path = BASE / utils.get_mode(sys.argv)
    data = utils.parse_input_lines(file_path)
    dim_y = len(data)
    dim_x = len(data[0])
    print("Part 1:", part1(data))
    print("Part 2:", part2(data))
