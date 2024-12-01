import { readFile } from '../shared'
import fs from 'fs';

const day = "20"

function SolvePart1(inputFile: string) {
    let numbers = readFile(inputFile).split("\n").map(l => +l)
    let l = numbers.length
    let indexes: number[] = []
    for (let i = 0; i < numbers.length; i++) {
        indexes.push(i)
    }

    let index = indexes.indexOf(6)
    let number = numbers[index]
    let newIndex = (index + numbers[index] + l + 1) % l

    let pN: number[] = []
    let nN: number[] = []
    let pI: number[] = []
    let nI: number[] = []
    if (newIndex > index) {
        for (let i = 0; i < newIndex; i++) {
            if (i != index) {
                pN.push(numbers[i])
                pI.push(indexes[i])
            }
        }
        for (let i = newIndex; i < l; i++) {
            nN.push(numbers[i])
            nI.push(indexes[i])
        }
    } else if (newIndex < index) {
        for (let i = 0; i < index; i++) {
            pN.push(numbers[i])
            pI.push(indexes[i])
        }
        for (let i = index; i < l; i++) {
            if (i != newIndex) {
                nN.push(numbers[i])
                nI.push(indexes[i])
            }
        }
    }
    numbers = [...pN, number, ...nN]
    indexes = [...pI, index, ...nI]

    console.log(numbers, indexes)

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