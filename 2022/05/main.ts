import { readFile } from '../shared'
import fs from 'fs';

const day = "05"

function ParseInput1(ls: string[]): [string[][], number[][]] {
    let splitInput = 0
    for (let i = 0; i < ls.length; i++) {
        if (ls[i] == "") {
            splitInput = i
        }
    }

    let stacks: string[][] = []
    for (let j = 1; j < ls[2].length; j += 4) {
        stacks.push([])
    }
    for (let i = splitInput - 2; i >= 0; i--) {
        for (let j = 0; j < stacks.length; j++) {
            let val = ls[i][j * 4 + 1]
            if (val != ' ' && val != undefined) {
                stacks[j].push(val)
            }
        }
    }

    let commands: number[][] = []
    for (let i = splitInput + 1; i < ls.length; i++) {
        commands.push([+ls[i].split(" ")[1], +ls[i].split(" ")[3], +ls[i].split(" ")[5]])
    }

    return [stacks, commands]
}

function ParseInput2(ls: string[]): [string[], number[][]] {
    let splitInput = 0
    for (let i = 0; i < ls.length; i++) {
        if (ls[i] == "") {
            splitInput = i
        }
    }

    let stacks: string[] = []
    for (let j = 1; j < ls[2].length; j += 4) {
        stacks.push("")
    }
    for (let i = splitInput - 2; i >= 0; i--) {
        for (let j = 0; j < stacks.length; j++) {
            let val = ls[i][j * 4 + 1]
            if (val != ' ' && val != undefined) {
                stacks[j] += val
            }
        }
    }

    let commands: number[][] = []
    for (let i = splitInput + 1; i < ls.length; i++) {
        commands.push([+ls[i].split(" ")[1], +ls[i].split(" ")[3], +ls[i].split(" ")[5]])
    }

    return [stacks, commands]
}

function MoveCrate(stacks: string[][], from: number, to: number): string[][] {
    let crateToMove = stacks[from - 1].pop()!
    stacks[to - 1].push(crateToMove)
    return stacks
}

function SolvePart1(inputFile: string) {
    let lines = readFile(inputFile).split("\n")
    let setup = ParseInput1(lines)

    for (let command = 0; command < setup[1].length; command++) {
        for (let i = 0; i < setup[1][command][0]; i++) {
            setup[0] = MoveCrate(setup[0], setup[1][command][1], setup[1][command][2])
        }
    }

    let tops = ""
    for (let i = 0; i < setup[0].length; i++) {
        tops += setup[0][i].pop()
    }

    return tops
}

function SolvePart2(inputFile: string) {
    let lines = readFile(inputFile).split("\n")
    let setup = ParseInput2(lines)

    for (let command = 0; command < setup[1].length; command++) {
        let nb = setup[1][command][0]
        let from = setup[1][command][1] - 1
        let to = setup[1][command][2] - 1

        let crates = setup[0][from].substring(setup[0][from].length - nb, setup[0][from].length)
        setup[0][from] = setup[0][from].substring(0, setup[0][from].length - nb)
        setup[0][to] += crates
    }

    let tops = ""
    for (let i = 0; i < setup[0].length; i++) {
        tops += setup[0][i][setup[0][i].length - 1]
    }

    return tops
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