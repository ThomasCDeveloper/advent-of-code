import sys
import utils
from pathlib import Path


def part1(data: list[str]):
    for i, line in enumerate(data):
        while "  " in line:
            line = line.replace("  ", " ")
        data[i] = line
    totals = list(map(int, data[0].split(" ")))
    operations = data[-1].split(" ")
    for line in data[1:-1]:
        numbers = list(map(int, line.split(" ")))
        for i, new in enumerate(numbers):
            if operations[i] == "+":
                totals[i] += numbers[i]
            else:
                totals[i] *= numbers[i]
    return sum(totals)


def part2(data: list[str]):
    total = 0

    rotated = []
    for i in range(len(data[0]) - 1):
        rotated.append("".join([data[j][i] for j in range(len(data))]))

    current_operation = ""
    subtotal = 0
    for i in range(len(rotated)):
        line = rotated[i]
        if line.strip() == "":
            total += subtotal
            continue
        if "+" in line:
            current_operation = "+"
            subtotal = int(line[:-1].strip())
            line = line[:-1]
        elif "*" in line:
            current_operation = "*"
            subtotal = int(line[:-1].strip())
            line = line[:-1]
        else:
            num = int(line.strip())
            if current_operation == "+":
                subtotal += num
            if current_operation == "*":
                subtotal *= num
    total += subtotal

    return total


if __name__ == "__main__":
    BASE = Path(__file__).resolve().parent
    file_path = BASE / utils.get_mode(sys.argv)
    data = utils.parse_raw_lines(file_path)
    print("Part 1:", part1([line.strip() for line in data]))
    print("Part 2:", part2(data))
