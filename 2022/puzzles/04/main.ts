import { readFile } from '../shared'
import fs from 'fs';

const day = "04"

function SolvePart1(inputFile: string) {
    let lines = readFile(inputFile).split("\n")

    let score = 0
    for (let i = 0; i < lines.length; i++) {
        const line = lines[i]

        let aa = +line.split(",")[0].split("-")[0]
        let ab = +line.split(",")[0].split("-")[1]
        let ba = +line.split(",")[1].split("-")[0]
        let bb = +line.split(",")[1].split("-")[1]

        if ((aa <= ba && ab >= bb) || (ba <= aa && bb >= ab)) {
            score++
        }
    }

    return score
}

function SolvePart2(inputFile: string) {
    let lines = readFile(inputFile).split("\n")

    let score = 0
    for (let i = 0; i < lines.length; i++) {
        const line = lines[i]

        let aa = +line.split(",")[0].split("-")[0]
        let ab = +line.split(",")[0].split("-")[1]
        let ba = +line.split(",")[1].split("-")[0]
        let bb = +line.split(",")[1].split("-")[1]

        if ((aa <= bb && aa >= ba) || (ab <= bb && ab >= ba) || (ba <= ab && ba >= aa)) {
            score++
        }
    }

    return score
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