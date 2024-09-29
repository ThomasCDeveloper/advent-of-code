import { readFile } from '../shared'
import fs from 'fs';

const day = "25"

function ConvertSNAFUToNumber(input: string): number {
    let sum = 0
    let mult = 1
    for (let i = input.length - 1; i >= 0; i--) {
        let digit = input[i]
        sum += ("=-012".indexOf(digit) - 2) * mult
        mult *= 5
    }
    return sum
}

function ConvertNumberToSNAFU(input: number): string {
    let output = ""

    while (input != 0) {
        let remainder = input % 5
        input = Math.floor(input / 5)

        output = "012=-"[remainder] + output
        if (remainder > 2) input += 1
    }

    return output
}

function SolvePart1(inputFile: string) {
    let compt = 0

    readFile(inputFile).split("\n").forEach((line) => {
        compt += ConvertSNAFUToNumber(line)
    })

    return ConvertNumberToSNAFU(compt)
}

console.log("== TEST ==")
console.log(" -Part 1: " + SolvePart1(day + "/a.example"))

if (true) {
    console.log("== DATA ==")
    console.log(" -Part 1: " + SolvePart1(day + "/ab.input"))
}