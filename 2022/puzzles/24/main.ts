import { readFile } from '../shared'
import fs from 'fs';

const day = "24"

const isEqual = (a: [number, number, number], b: [number, number, number]) =>
    a[0] === b[0] && a[1] === b[1] && a[2] === b[2]

function SolvePart1(inputFile: string) {
    let lines = readFile(inputFile).split("\n")

    let blizzards: [number, number][][] = [[], [], [], []] // ><^v

    for (let y = 0; y < lines.length; y++) {
        for (let x = 0; x < lines[y].length; x++) {
            let index = "><^v".indexOf(lines[y][x])
            if (index >= 0) {
                blizzards[index].push([x, y])
            }
        }
    }

    let queue: [number, number, number][] = [[1, 0, 0]] // x, y, time
    let visited: [number, number, number][] = []

    let i = 0

    while (queue.length != 0) {
        let status = queue.shift()!

        for (let [x, y] of [[status[0], status[1]], [status[0] + 1, status[1]], [status[0] - 1, status[1]], [status[0], status[1] + 1], [status[0], status[1] - 1]]) {
            if (x < 1 || x >= lines[0].length - 1 || y < 0 || y >= lines.length) {
                continue
            }
            if (lines[y][x] == "#") {
                continue
            }

            if (x == lines[0].length - 2 && y == lines.length - 1) {
                return status[2] + 1
            }

            let nextStatus: [number, number, number] = [x, y, status[2] + 1]

            if (visited.some(v => isEqual(v, nextStatus)) || queue.some(q => isEqual(q, nextStatus))) {
                continue
            }

            // left to check if player is on a blizzard
            // think of a way to move player instead of 654615654 blizzards

            queue.push(nextStatus)
            visited.push(nextStatus)
        }


        queue.sort((a, b) => a[2] - b[2])
        /*if (i == 8) return 0
        i++*/
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

if (false) {
    console.log("== DATA ==")
    console.log(" -Part 1: " + SolvePart1(day + "/ab.input"))
    console.log(" -Part 2: " + SolvePart2(day + "/ab.input"))
}