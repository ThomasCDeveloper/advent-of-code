import { readFile } from '../shared'
import fs from 'fs';

const day = "08"

function SolvePart1(inputFile: string) {
    let lines = readFile(inputFile).split("\n")

    let areVisible: boolean[][] = []
    for (let row = 0; row < lines.length; row++) {
        areVisible.push([])
        for (let col = 0; col < lines[0].length; col++) {
            areVisible[row].push(false)

            let visibleLeft = true
            for (let i = 0; i < col; i++) {
                if (lines[row][i] >= lines[row][col]) {
                    visibleLeft = false
                    break
                }
            }
            let visibleRight = true
            for (let i = lines[0].length - 1; i > col; i--) {
                if (lines[row][i] >= lines[row][col]) {
                    visibleRight = false
                    break
                }
            }
            let visibleUp = true
            for (let i = 0; i < row; i++) {
                if (lines[i][col] >= lines[row][col]) {
                    visibleUp = false
                    break
                }
            }
            let visibleDown = true
            for (let i = lines.length - 1; i > row; i--) {
                if (lines[i][col] >= lines[row][col]) {
                    visibleDown = false
                    break
                }
            }


            areVisible[row][col] = visibleLeft || visibleRight || visibleUp || visibleDown
        }
    }

    let total = 0
    for (let row = 0; row < lines.length; row++) {
        for (let col = 0; col < lines[0].length; col++) {
            if (areVisible[row][col]) {
                total++
            }
        }
    }

    return total
}

function SolvePart2(inputFile: string) {
    let lines = readFile(inputFile).split("\n")

    let max = 0
    for (let row = 0; row < lines.length; row++) {
        for (let col = 0; col < lines[0].length; col++) {
            let valueLeft = 0
            for (let i = 1; i < col + 1; i++) {
                valueLeft++
                if (lines[row][col - i] >= lines[row][col]) {
                    break
                }
            }
            let valueRight = 0
            for (let i = 1; i < lines[0].length - col; i++) {
                valueRight++
                if (lines[row][col + i] >= lines[row][col]) {
                    break
                }
            }
            let valueUp = 0
            for (let i = 1; i < row + 1; i++) {
                valueUp++
                if (lines[row - i][col] >= lines[row][col]) {
                    break
                }
            }
            let valueDown = 0
            for (let i = 1; i < lines.length - row; i++) {
                valueDown++
                if (lines[row + i][col] >= lines[row][col]) {
                    break
                }
            }

            let mult = valueLeft * valueRight * valueUp * valueDown
            if (mult > max) {
                max = mult
            }
        }
    }
    return max
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