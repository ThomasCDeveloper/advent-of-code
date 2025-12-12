import sys
import utils
from pathlib import Path


def part1(data):
    shapes_str = data.split("\n\n")[:-1]
    for s in shapes_str:
        shape = []
        for y, li in enumerate(s.split("\n")[1:]):
            for x, co in enumerate(s.split("\n")[y + 1]):
                if co == "#":
                    shape.append([x, y])
        shapes.append(shape)
    regions_str = data.split("\n\n")[-1]
    regions = []
    for r in regions_str.split("\n"):
        size = list(map(int, r.split(":")[0].split("x")))
        nb_items = list(map(int, r.split(": ")[1].split(" ")))
        regions.append([size, nb_items])

    total_possible = 0
    for region in regions:
        area = region[0][0] * region[0][1]
        if area >= sum(region[1]) * 7:
            total_possible += 1

    return total_possible


shapes = []
if __name__ == "__main__":
    BASE = Path(__file__).resolve().parent
    file_path = BASE / utils.get_mode(sys.argv)
    data = utils.parse_input(file_path)
    print("Part 1:", part1(data))
