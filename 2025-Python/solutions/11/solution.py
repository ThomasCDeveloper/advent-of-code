import sys
import utils
from pathlib import Path
import functools


@functools.lru_cache()
def count_paths(label):
    total = 0
    for next_label in labels[label]:
        if next_label == "out":
            return 1
        total += count_paths(next_label)
    return total


def part1():
    return count_paths("you")


@functools.lru_cache()
def count_paths_fft_dac(label):
    if label == "out":
        return (0, 0, 0, 1)

    total_both = 0
    total_fft_only = 0
    total_dac_only = 0
    total_none = 0

    for next_label in labels[label]:
        b, f, d, n = count_paths_fft_dac(next_label)

        if label == "fft":
            total_both += b
            total_fft_only += f + n
            total_both += d
        elif label == "dac":
            total_both += b
            total_dac_only += d + n
            total_both += f
        else:
            total_both += b
            total_fft_only += f
            total_dac_only += d
            total_none += n

    return (total_both, total_fft_only, total_dac_only, total_none)


def part2():
    a, _, _, _ = count_paths_fft_dac("svr")
    return a


labels = {}
if __name__ == "__main__":
    BASE = Path(__file__).resolve().parent
    file_path = BASE / utils.get_mode(sys.argv)
    data = utils.parse_input_lines(file_path)
    for line in data:
        labels[line.split(": ")[0]] = line.split(": ")[1].split(" ")
    print("Part 1:", part1())
    print("Part 2:", part2())
