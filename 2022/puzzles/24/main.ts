import { readFile } from '../shared'
import fs from 'fs';

const day = "24"

const isStatusSame = (a: [number, number, number], b: [number, number, number]) =>
    a[0] === b[0] && a[1] === b[1] && a[2] === b[2]

const isPositionSame = (a: [number, number], b: [number, number]) =>
    a[0] === b[0] && a[1] === b[1]

const gcd = (a: number, b: number): number => {
    const r = a % b

    if (r === 0) {
        return b
    }

    return gcd(b, a % b)
}

function SolvePart1(inputFile: string) {
    let lines = readFile(inputFile).split("\n")
    let map: string[][] = []
    for (let y = 0; y < lines.length - 2; y++) {
        map.push([])
        for (let x = 1; x < lines[y + 1].length - 1; x++) {
            map[y].push(lines[y + 1][x])
        }
    }

    let loop = map.length * map[0].length / gcd(map.length, map[0].length)

    let directions = [
        [map[0].length - 1, 0],
        [1, 0],
        [0, 1],
        [0, map.length - 1],
    ]

    let blizzards: [number, number][][] = [[], [], [], []] // ><^v

    for (let y = 0; y < map.length; y++) {
        for (let x = 0; x < map[y].length; x++) {
            let index = "><^v".indexOf(map[y][x])
            if (index >= 0) {
                blizzards[index].push([x, y])
            }
        }
    }

    let queue: [number, number, number][] = [[0, -1, 0]] // x, y, time
    let visited: [number, number, number][] = []

    while (queue.length != 0) {
        let status = queue.shift()!

        for (let [x, y] of [[status[0], status[1]], [status[0] + 1, status[1]], [status[0] - 1, status[1]], [status[0], status[1] + 1], [status[0], status[1] - 1]]) {
            if (x == map[0].length - 1 && y == map.length) {
                return status[2]
            }
            if ((x < 0 || x >= map[0].length || y < 0 || y >= map.length) && (x != 0 || y != -1)) {
                continue
            }

            let nextStatus: [number, number, number] = [x, y, (status[2] + 1)]

            if (visited.some(v => isStatusSame(v, [nextStatus[0], nextStatus[1], nextStatus[2] % loop])) || queue.some(q => isStatusSame(q, [nextStatus[0], nextStatus[1], nextStatus[2] % loop]))) {
                continue
            }

            // left to check if player is on a blizzard
            // think of a way to move player instead of 654615654 blizzards
            let isOnBlizzard = false
            for (let i = 0; i < 4; i++) {
                if (blizzards[i].some(b => isPositionSame(b, [(x + status[2] * directions[i][0]) % map[0].length, (y + status[2] * directions[i][1]) % map.length]))) {
                    isOnBlizzard = true
                    break
                }
            }
            if (isOnBlizzard) {
                continue
            }

            queue.push(nextStatus)
            visited.push([nextStatus[0], nextStatus[1], nextStatus[2] % loop])
        }


        queue.sort((a, b) => a[2] - b[2])
        console.log(queue.length)
    }

    return 0
}

function SolvePart2(inputFile: string) {

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