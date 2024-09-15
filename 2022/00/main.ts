import { readFile } from '../shared'
import fs from 'fs';

const day = "00

function SolvePart1(inputFile: string) {

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

console.log("== DATA ==")
console.log(" -Part 1: " + SolvePart1(day + "/ab.input"))
console.log(" -Part 2: " + SolvePart2(day + "/ab.input"))