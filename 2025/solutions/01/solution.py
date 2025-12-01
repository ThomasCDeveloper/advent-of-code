import sys


def parse_input(file_path):
    with open(file_path) as f:
        return [line.strip() for line in f.readlines()]


def part1(data):
    position = 50
    count = 0

    for command in data:
        direction = command[0]
        amount = int(command[1:])
        if direction == "R":
            amount = 100 - amount
        position = (position + amount) % 100
        if position == 0:
            count += 1

    return count


def part2(data):
    position = 50
    last = 50
    count = 0

    for command in data:
        direction = command[0]
        amount = int(command[1:])

        count += amount // 100
        amount %= 100
        if amount == 0:
            continue

        if direction == "R":
            position = position + amount
        else:
            position = position - amount

        if (position >= 100 and direction == "R") or (
            direction == "L" and position < 1 and last > 0
        ):
            count += 1

        position = (position + 100) % 100

        last = position

    return count


if __name__ == "__main__":
    if len(sys.argv) != 2 or sys.argv[1] not in ("input", "test"):
        sys.exit(1)

    file_path = (
        "./solutions/01/input.txt"
        if sys.argv[1] == "input"
        else "./solutions/01/test.txt"
    )

    data = parse_input(file_path)
    print("Part 1:", part1(data))
    print("Part 2:", part2(data))
