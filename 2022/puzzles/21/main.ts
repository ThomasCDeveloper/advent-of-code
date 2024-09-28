import { readFile } from '../shared'
import fs from 'fs';

const day = "21"

let monkeyNumbers: Record<string, number> = {}
let monkeyRules: Record<string, string[]> = {}

function GetNumber(monkey: string): number {
    if (monkey in monkeyNumbers) {
        return monkeyNumbers[monkey]
    }

    let number = 0

    switch (monkeyRules[monkey][1]) {
        case "+":
            number = GetNumber(monkeyRules[monkey][0]) + GetNumber(monkeyRules[monkey][2])
            break
        case "-":
            number = GetNumber(monkeyRules[monkey][0]) - GetNumber(monkeyRules[monkey][2])
            break
        case "*":
            number = GetNumber(monkeyRules[monkey][0]) * GetNumber(monkeyRules[monkey][2])
            break
        default:
            number = GetNumber(monkeyRules[monkey][0]) / GetNumber(monkeyRules[monkey][2])
            break
    }

    monkeyNumbers[monkey] = number

    return number
}

function SolvePart1(inputFile: string) {
    monkeyNumbers = {}

    readFile(inputFile).split("\n").forEach((line) => {
        let monkey = line.split(": ")[0]
        let rule = line.split(": ")[1]
        if (rule.split(" ").length == 1) {
            monkeyNumbers[monkey] = +rule
        } else {
            monkeyRules[monkey] = rule.split(" ")
        }
    })

    return GetNumber("root")
}

function SolvePart2(inputFile: string) {
    let leftValue: number[] = []
    let rightValue: number[] = []

    for (let humn of [0, 1]) {
        monkeyNumbers = {}

        readFile(inputFile).split("\n").forEach((line) => {
            let monkey = line.split(": ")[0]
            let rule = line.split(": ")[1]
            if (rule.split(" ").length == 1) {
                monkeyNumbers[monkey] = +rule
            } else {
                monkeyRules[monkey] = rule.split(" ")
            }
        })

        monkeyNumbers["humn"] = humn

        leftValue.push(GetNumber(monkeyRules["root"][0]))
        rightValue.push(GetNumber(monkeyRules["root"][2]))
    }

    return (rightValue[0] - leftValue[0]) / (leftValue[1] - leftValue[0])
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