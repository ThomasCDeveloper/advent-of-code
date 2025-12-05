import { readFile } from '../shared'
import fs from 'fs';

const day = "06"

function SolvePart1(inputFile: string) {
    let input = readFile(inputFile)

    for (let i = 0; i < input.length - 4; i++) {
        let subString = input.substring(i, i + 4)
        let isKey = true
        for (let c = 0; c < 3; c++) {
            for (let c2 = c + 1; c2 < 4; c2++) {
                if (subString[c] == subString[c2]) {
                    isKey = false
                }
            }
        }
        if (isKey) {
            return i + 4
        }
    }

    return -1
}

function SolvePart2(inputFile: string) {
    let input = readFile(inputFile)

    for (let i = 0; i < input.length - 14; i++) {
        let subString = input.substring(i, i + 14)
        let isKey = true
        for (let c = 0; c < 13; c++) {
            for (let c2 = c + 1; c2 < 14; c2++) {
                if (subString[c] == subString[c2]) {
                    isKey = false
                }
            }
        }
        if (isKey) {
            return i + 14
        }
    }

    return -1
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