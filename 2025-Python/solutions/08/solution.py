import sys
import utils
from pathlib import Path
import heapq
import numpy as np


def distance(c_a, c_b):
    return (c_a[0] - c_b[0]) ** 2 + (c_a[1] - c_b[1]) ** 2 + (c_a[2] - c_b[2]) ** 2


class UnionFind:
    def __init__(self, n):
        self.parent = list(range(n))
        self.rank = [0] * n
        self.size = [1] * n

    def find(self, x):
        while self.parent[x] != x:
            self.parent[x] = self.parent[self.parent[x]]
            x = self.parent[x]
        return x

    def union(self, a, b):
        ra = self.find(a)
        rb = self.find(b)
        if ra == rb:
            return ra

        if self.rank[ra] < self.rank[rb]:
            self.parent[ra] = rb
            self.size[rb] += self.size[ra]
            return rb
        elif self.rank[rb] < self.rank[ra]:
            self.parent[rb] = ra
            self.size[ra] += self.size[rb]
            return ra
        else:
            self.parent[rb] = ra
            self.rank[ra] += 1
            self.size[ra] += self.size[rb]
            return ra


def part1(cubes, pairs, n, iter):
    smallest = heapq.nsmallest(iter, pairs, key=lambda p: p[0])

    uf = UnionFind(n)
    for _, a, b in smallest:
        uf.union(a, b)

    group_map = {}
    for i in range(n):
        r = uf.find(i)
        group_map.setdefault(r, 0)
        group_map[r] += 1

    sizes = sorted(group_map.values(), reverse=True)
    return sizes[0] * sizes[1] * sizes[2]


def part2(cubes, pairs, n):
    pairs.sort(key=lambda p: p[0])

    uf = UnionFind(n)

    for _, a, b in pairs:
        root = uf.union(a, b)

        if uf.size[root] == n:
            return cubes[a][0] * cubes[b][0]

    return "not found"


if __name__ == "__main__":
    BASE = Path(__file__).resolve().parent
    mode = utils.get_mode(sys.argv)
    file_path = BASE / mode
    data = utils.parse_input_lines(file_path)
    cubes = [list(map(int, x.split(","))) for x in data]
    n = len(cubes)
    pairs = [
        (distance(cubes[a], cubes[b]), a, b) for a in range(n) for b in range(a + 1, n)
    ]
    print(
        "Part 1:",
        part1(cubes, pairs, n, 10 if mode == "test.txt" else 1000),
    )
    print("Part 2:", part2(cubes, pairs, n))
