import sys
import utils
from pathlib import Path


def distance(c_a, c_b):
    return (c_a[0] - c_b[0]) ** 2 + (c_a[1] - c_b[1]) ** 2 + (c_a[2] - c_b[2]) ** 2


def part1(cubes, iter):
    couples = []
    for a, c_a in enumerate(cubes[:-1]):
        for b, c_b in enumerate(cubes[a + 1 :]):
            couples.append([c_a, c_b])

    pairs = sorted(couples, key=lambda x: distance(x[0], x[1]))[:iter]

    groups = []

    for a, b in pairs:
        matching = []
        for g in groups:
            if any(x == a or x == b for x in g):
                matching.append(g)

        if not matching:
            groups.append([a, b])
        else:
            merged = []
            for g in matching:
                merged.extend(g)
                groups.remove(g)

            merged.extend([a, b])

            final = []
            for x in merged:
                if not any(x == y for y in final):
                    final.append(x)

            groups.append(final)

    sizes = sorted(list(map(len, groups)), reverse=True)
    return sizes[0] * sizes[1] * sizes[2]


def part2(cubes):
    couples = []
    for a, c_a in enumerate(cubes[:-1]):
        for b, c_b in enumerate(cubes[a + 1 :]):
            couples.append([c_a, c_b])

    pairs = sorted(couples, key=lambda x: distance(x[0], x[1]))

    groups = []
    for a, b in pairs:
        matching = []
        for g in groups:
            if any(x == a or x == b for x in g):
                matching.append(g)

        if not matching:
            groups.append([a, b])
        else:
            merged = []
            for g in matching:
                merged.extend(g)
                groups.remove(g)

            merged.extend([a, b])

            final = []
            for x in merged:
                if not any(x == y for y in final):
                    final.append(x)

            groups.append(final)

        sizes = sorted(list(map(len, groups)), reverse=True)
        if sizes[0] == len(data):
            return a[0] * b[0]

    return "not found"


if __name__ == "__main__":
    BASE = Path(__file__).resolve().parent
    mode = utils.get_mode(sys.argv)
    file_path = BASE / mode
    data = utils.parse_input_lines(file_path)
    cubes = [list(map(int, x.split(","))) for x in data]
    print(
        "Part 1:",
        part1(cubes, 10 if mode == "test.txt" else 1000),
    )
    print("Part 2:", part2(cubes))
