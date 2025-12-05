import { readFile } from '../shared'
import fs from 'fs';

const day = "24"

const gcd = (a: number, b: number): number => {
    const r = a % b

    if (r === 0) {
        return b
    }

    return gcd(b, a % b)
}

function SolvePart1(inputFile: string) {
    let lines = readFile(inputFile).split("\n")
    let blizzards: Set<string>[] = Array.from({ length: 4 }, () => new Set()) // ><^v

    for (let y = 1; y < lines.length; y++) {
        for (let x = 1; x < lines[y].length; x++) {
            let index = "><^v".indexOf(lines[y][x])
            if (index >= 0) {
                blizzards[index].add(`${x - 1},${y - 1}`)
            }
        }
    }

    const [mx, my] = [lines[0].length - 2, lines.length - 2]

    let loop = mx * my / gcd(mx, my)

    let directions = [
        [mx - 1, 0],
        [1, 0],
        [0, 1],
        [0, my - 1],
    ]

    let queue: [number, number, number][] = [[0, -1, 0]] // x, y, time
    let visited = new Set<string>()

    while (queue.length != 0) {
        let status = queue.shift()!

        for (let [dx, dy] of [[0, 0], [1, 0], [-1, 0], [0, -1], [0, 1]]) {
            const [x, y] = [status[0] + dx, status[1] + dy]

            if (x == mx - 1 && y == my) {
                return status[2]
            }
            if ((x < 0 || x >= mx || y < 0 || y >= my) && (x != 0 || y != -1)) {
                continue
            }

            let isOnBlizzard = false
            for (let i = 0; i < 4; i++) {
                if (blizzards[i].has(`${(x + status[2] * directions[i][0]) % mx},${(y + status[2] * directions[i][1]) % my}`)) {
                    isOnBlizzard = true
                    break
                }
            }
            if (isOnBlizzard) {
                continue
            }

            let nextStatus: [number, number, number] = [x, y, status[2] + 1]
            const key = `${x},${y},${nextStatus[2] % loop}`

            if (visited.has(key)) {
                continue
            }

            queue.push(nextStatus)
            visited.add(key)
        }

        queue.sort((a, b) => a[2] - b[2])
    }
    return 0
}

function SolvePart2(inputFile: string) {
    let lines = readFile(inputFile).split("\n")
    let blizzards: Set<string>[] = Array.from({ length: 4 }, () => new Set()) // ><^v

    for (let y = 1; y < lines.length; y++) {
        for (let x = 1; x < lines[y].length; x++) {
            let index = "><^v".indexOf(lines[y][x])
            if (index >= 0) {
                blizzards[index].add(`${x - 1},${y - 1}`)
            }
        }
    }

    const [mx, my] = [lines[0].length - 2, lines.length - 2]

    let loop = mx * my / gcd(mx, my)

    let directions = [
        [mx - 1, 0],
        [1, 0],
        [0, 1],
        [0, my - 1],
    ]

    let targets = [
        [mx - 1, my],
        [0, -1]
    ]

    let queue: [number, number, number, number][] = [[0, -1, 0, 0]] // x, y, time, step
    let visited = new Set<string>()

    while (queue.length != 0) {
        let [ox, oy, time, step] = queue.shift()!
        time += 1

        for (let [dx, dy] of [[0, 0], [1, 0], [-1, 0], [0, -1], [0, 1]]) {
            const [x, y] = [ox + dx, oy + dy]

            let nstep = step
            if (x == targets[step % 2][0] && y == targets[step % 2][1]) {
                if (step == 2)
                    return time
                nstep++
            }
            if ((x < 0 || x >= mx || y < 0 || y >= my) && ((x != 0 || y != -1) && (x != mx - 1 || y != my))) {
                continue
            }

            let isOnBlizzard = false
            if (((x != 0 || y != -1) && (x != mx - 1 || y != my)))
                for (let i = 0; i < 4; i++) {
                    if (blizzards[i].has(`${(x + time * directions[i][0]) % mx},${(y + time * directions[i][1]) % my}`)) {
                        isOnBlizzard = true
                        break
                    }
                }
            if (isOnBlizzard) {
                continue
            }

            const key = `${x},${y},${time % loop},${step}`

            if (visited.has(key)) {
                continue
            }

            queue.push([x, y, time, nstep])
            visited.add(key)
        }

        queue.sort((a, b) => a[2] - b[2])
    }
    return 0
}

console.log("== TEST ==")
console.log(" -Part 1: " + SolvePart1(day + "/a.example"))
if (fs.existsSync(day + "/b.example")) {
    console.log(" -Part 2: " + SolvePart2(day + "/b.example"))
} else {
    console.log(" -Part 2: " + SolvePart2(day + "/a.example"))
}

if (true) {
    console.log("== DATA ==")
    console.log(" -Part 1: " + SolvePart1(day + "/ab.input"))
    console.log(" -Part 2: " + SolvePart2(day + "/ab.input"))
}