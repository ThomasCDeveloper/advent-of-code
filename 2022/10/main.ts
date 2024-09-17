import { readFile } from '../shared'
import fs from 'fs';

const day = "10"

function SolvePart1(inputFile: string) {
    let commands = readFile(inputFile).split("\n")

    let X = 1
    let cycle = 0
    let sum = 0
    for (let i = 0; i < commands.length; i++) {
        if (commands[i] == "noop") {
            cycle++
            if (cycle % 40 == 20) {
                sum += cycle * X
            }
        } else {
            for (let j = 0; j < 2; j++) {
                cycle++
                if (cycle % 40 == 20) {
                    sum += cycle * X
                }
            }
            X += +(commands[i].split(" ")[1])
        }
    }

    return sum
}

function Print(screen: string[]) {
    for (let row = 0; row < 6; row++) {
        let l = ""
        for (let col = 0; col < 40; col++) {
            l += screen[row * 40 + col]
        }
        console.log(l)
    }
}

function SolvePart2(inputFile: string) {
    let screen: string[] = []
    for (let i = 0; i < 6 * 40; i++) {
        screen.push(".")
    }

    let commands = readFile(inputFile).split("\n")

    let X = 1
    let cycle = 0
    for (let i = 0; i < commands.length; i++) {
        if (commands[i] == "noop") {
            if (Math.abs(cycle % 40 - X) <= 1) {
                screen[cycle] = "#"
            }
            cycle++
        } else {
            for (let j = 0; j < 2; j++) {
                if (Math.abs(cycle % 40 - X) <= 1) {
                    screen[cycle] = "#"
                }
                cycle++
            }
            X += +(commands[i].split(" ")[1])
        }
    }

    Print(screen)

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