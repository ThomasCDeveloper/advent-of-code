import { readFile } from '../shared'
import fs from 'fs';

const day = "03"

function SolvePart1(inputFile: string) {
    let score = 0
    let lines = readFile(inputFile).split("\n")

    for (let i = 0; i < lines.length; i++) {
        let pocket0 = lines[i].substring(0, lines[i].length / 2)
        let pocket1 = lines[i].substring(lines[i].length / 2, lines[i].length)

        for (let numChar = 0; numChar < pocket0.length; numChar++) {
            let char = pocket0[numChar]
            if (pocket1.indexOf(char) >= 0) {
                score += "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ".indexOf(char) + 1
                break
            }
        }
    }

    return score
}

function SolvePart2(inputFile: string) {
    let score = 0
    let lines = readFile(inputFile).split("\n")

    for (let i = 0; i < lines.length; i += 3) {
        let elf0 = lines[i]
        let elf1 = lines[i + 1]
        let elf2 = lines[i + 2]

        for (let numChar = 0; numChar < elf0.length; numChar++) {
            let char = elf0[numChar]
            if (elf1.indexOf(char) >= 0 && elf2.indexOf(char) >= 0) {
                score += "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ".indexOf(char) + 1
                break
            }
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