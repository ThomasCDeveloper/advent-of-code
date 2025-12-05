import sys
import utils
from pathlib import Path


def part1(data):
    return 0  # TODO


def part2(data):
    return 0  # TODO


if __name__ == "__main__":
    BASE = Path(__file__).resolve().parent
    file_path = BASE / utils.get_mode(sys.argv)
    data = utils.parse_input(file_path)
    print("Part 1:", part1(data))
    print("Part 2:", part2(data))
