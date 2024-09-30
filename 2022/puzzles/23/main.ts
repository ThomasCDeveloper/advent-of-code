import { readFile } from '../shared'
import fs from 'fs';

const day = "23"

function SolvePart1(inputFile: string) {
    let monkeys: [number, number][] = []

    let lines = readFile(inputFile).split("\n")
    for (let y = 0; y < lines.length; y++) {
        for (let x = 0; x < lines[y].length; x++) {
            if (lines[y][x] == "#") {
                monkeys.push([x, y])
            }
        }
    }

    for (let i = 0; i < 10; i++) {
        // do stuff
    }

    let maxX = Math.max(...monkeys.map((m) => m[0]))
    let maxY = Math.max(...monkeys.map((m) => m[1]))
    let minX = Math.min(...monkeys.map((m) => m[0]))
    let minY = Math.min(...monkeys.map((m) => m[1]))

    return (maxX - minX + 1) * (maxY - minY + 1) - monkeys.length
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