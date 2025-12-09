import sys
import utils
from pathlib import Path


def area(p1, p2):
    return (abs(p1[0] - p2[0]) + 1) * (abs(p1[1] - p2[1]) + 1)


def part1(corners):
    max_area = 0
    for corner1 in corners:
        for corner2 in corners:
            area_12 = area(corner1, corner2)
            if area_12 > max_area:
                max_area = area_12
    return max_area


def segment_hits_interior_rect(seg, xmin, xmax, ymin, ymax):
    (ax, ay), (bx, by) = seg
    if ax == bx:
        x = ax
        if not (xmin < x < xmax):
            return False
        sy0, sy1 = sorted([ay, by])
        return max(sy0, ymin) < min(sy1, ymax)

    y = ay
    if not (ymin < y < ymax):
        return False
    sx0, sx1 = sorted([ax, bx])
    return max(sx0, xmin) < min(sx1, xmax)


def part2(corners):
    max_area = 0
    n = len(corners)
    segments = [[corners[i], corners[(i + 1) % n]] for i in range(n)]

    for i in range(n - 1):
        c1 = corners[i]
        for j in range(i + 1, n):
            c2 = corners[j]
            area_12 = area(c1, c2)
            if area_12 <= max_area:
                continue

            xmin, xmax = sorted([c1[0], c2[0]])
            ymin, ymax = sorted([c1[1], c2[1]])

            is_valid = True
            for seg in segments:
                if segment_hits_interior_rect(seg, xmin, xmax, ymin, ymax):
                    is_valid = False
                    break
            if is_valid:
                max_area = area_12

    return max_area


if __name__ == "__main__":
    BASE = Path(__file__).resolve().parent
    file_path = BASE / utils.get_mode(sys.argv)
    data = utils.parse_input_lines(file_path)
    corners = [list(map(int, x.split(","))) for x in data]
    print("Part 1:", part1(corners))
    print("Part 2:", part2(corners))
