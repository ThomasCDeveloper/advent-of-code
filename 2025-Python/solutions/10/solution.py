import sys
import utils
from pathlib import Path
from collections import deque


class Machine:
    def __init__(self, line: str):
        self.get_config(line)
        self.len = len(self.config)
        self.get_buttons(line)
        self.get_joltage(line)

    def get_config(self, line: str):
        self.config = [0 if x == "." else 1 for x in line.split("]")[0][1:]]

    def get_buttons(self, line: str):
        buttons = [x[1:-1] for x in line.split("] ")[1].split(" {")[0].split(" ")]
        self.buttons = []
        for b in buttons:
            c = [0] * self.len
            for n in list(map(int, b.split(","))):
                c[n] = 1
            self.buttons.append(c)

    def get_joltage(self, line: str):
        self.joltage = list(map(int, line.split("{")[1][:-1].split(",")))


def merge_actions(actions):
    a = [0] * len(actions[0])
    for action in actions:
        for i in range(len(action)):
            a[i] = (a[i] + action[i]) % 2
    return a


def xor(a, b):
    return [(a[i] ^ b[i]) for i in range(len(a))]


def part1(machines: list[Machine]) -> int:
    total = 0
    for m in machines:
        start = tuple([0] * m.len)
        target = tuple(m.config)
        q = deque([(start, 0)])
        visited = {start}

        found = False
        while q and not found:
            state, dist = q.popleft()
            for btn in m.buttons:
                new_state = tuple(xor(list(state), btn))
                if new_state == target:
                    total += dist + 1
                    found = True
                    break
                if new_state not in visited:
                    visited.add(new_state)
                    q.append((new_state, dist + 1))

    return total


def add(a, b):
    return [(a[i] + b[i]) for i in range(len(a))]


def gauss_rational(A, b):
    A = [row[:] for row in A]
    b = b[:]
    M = len(A)
    N = len(A[0])

    col = 0
    row = 0
    pivot_cols = []

    while row < M and col < N:
        pivot = row
        while pivot < M and A[pivot][col] == 0:
            pivot += 1
        if pivot == M:
            col += 1
            continue

        A[row], A[pivot] = A[pivot], A[row]
        b[row], b[pivot] = b[pivot], b[row]

        pivot_val = A[row][col]
        A[row] = [x / pivot_val for x in A[row]]
        b[row] /= pivot_val

        for r in range(M):
            if r != row and A[r][col] != 0:
                factor = A[r][col]
                A[r] = [A[r][c] - factor * A[row][c] for c in range(N)]
                b[r] -= factor * b[row]

        pivot_cols.append(col)
        row += 1
        col += 1

    return A, b, pivot_cols


def solve_integer_linear(A, b):
    N = len(A[0])

    Ar, br, pivots = gauss_rational(A, b)

    free_vars = [i for i in range(N) if i not in pivots]

    max_val = max(b) if b else 0
    limit = max_val + 5

    best = None
    best_sum = 10**18

    def backtrack(idx, x):
        nonlocal best, best_sum

        if idx == len(free_vars):
            sol = x[:]
            sol = sol[:]

            for r, col in enumerate(pivots):
                val = br[r]
                for j in range(N):
                    if j in free_vars:
                        val -= Ar[r][j] * sol[j]
                if abs(float(val) - round(val)) > 1e-9:
                    return
                sol[col] = int(round(val))

            if any(v < 0 for v in sol):
                return

            total = sum(sol)
            if total < best_sum:
                best_sum = total
                best = sol[:]
            return

        j = free_vars[idx]
        for v in range(0, limit):
            x[j] = v
            backtrack(idx + 1, x)

    x0 = [0] * N
    backtrack(0, x0)
    return best


def part2(machines):
    total = 0
    for i, m in enumerate(machines):
        A = m.buttons
        A = list(map(list, zip(*A)))
        b = m.joltage

        sol = solve_integer_linear(A, b)
        total += sum(sol)

    return total


if __name__ == "__main__":
    BASE = Path(__file__).resolve().parent
    file_path = BASE / utils.get_mode(sys.argv)
    data = utils.parse_input_lines(file_path)
    machines = []
    for line in data:
        machines.append(Machine(line))
    print("Part 1:", part1(machines))
    print("Part 2:", part2(machines))
