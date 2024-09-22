import { readFile } from '../shared'
import fs from 'fs';

const day = "11"

class Monkey {
    items: number[]
    operation: Function
    test: Function

    constructor(items: number[], operation: Function, test: Function, divider: number) {
        this.items = items
        this.operation = (old: number) => divider == 0 ? Math.floor(operation(old) / 3) : operation(old) % divider
        this.test = test
    }
}

function SolvePart1(inputFile: string) {
    let lines = readFile(inputFile).split("\n")
    let monkeys: Monkey[] = []

    for (let l = 0; l < lines.length; l += 7) {
        let items: number[] = []
        for (let n = 0; n < lines[l + 1].split(": ")[1].split(", ").length; n++) {
            items.push(+lines[l + 1].split(": ")[1].split(", ")[n])
        }
        let operation: Function
        let rule = lines[l + 2].split("= old ")[1]
        if (rule[0] == "+") {
            operation = (old: number) => old + +(rule.split(" ")[1])
        } else {
            if (rule[2] == "o") {
                operation = (old: number) => old * old
            } else {
                operation = (old: number) => old * +(rule.split(" ")[1])
            }
        }
        let test = (old: number) => (old % +lines[l + 3].split("by ")[1] == 0 ? +lines[l + 4].split("monkey ")[1] : +lines[l + 5].split("monkey ")[1])

        monkeys.push(new Monkey(items, operation, test, 0))
    }

    let counts: number[] = []
    for (let i = 0; i < monkeys.length; i++) {
        counts.push(0)
    }

    for (let round = 0; round < 20; round++) {
        for (let i = 0; i < monkeys.length; i++) {
            for (let item = 0; item < monkeys[i].items.length; item++) {
                counts[i]++
                let worry = monkeys[i].operation(monkeys[i].items[item])
                let nextMonkey = monkeys[i].test(worry)
                monkeys[nextMonkey].items.push(worry)
            }
            monkeys[i].items = []
        }
    }

    counts = counts.sort((a, b) => b - a)

    return counts[0] * counts[1]
}

function SolvePart2(inputFile: string) {
    let lines = readFile(inputFile).split("\n")
    let monkeys: Monkey[] = []

    let divider = 1
    for (let l = 0; l < lines.length; l += 7) {
        divider *= +lines[l + 3].split("by ")[1]
    }


    for (let l = 0; l < lines.length; l += 7) {
        let items: number[] = []
        for (let n = 0; n < lines[l + 1].split(": ")[1].split(", ").length; n++) {
            items.push(+lines[l + 1].split(": ")[1].split(", ")[n])
        }
        let operation: Function
        let rule = lines[l + 2].split("= old ")[1]
        if (rule[0] == "+") {
            operation = (old: number) => old + +(rule.split(" ")[1])
        } else {
            if (rule[2] == "o") {
                operation = (old: number) => old * old
            } else {
                operation = (old: number) => old * +(rule.split(" ")[1])
            }
        }
        let test = (old: number) => (old % +lines[l + 3].split("by ")[1] == 0 ? +lines[l + 4].split("monkey ")[1] : +lines[l + 5].split("monkey ")[1])

        monkeys.push(new Monkey(items, operation, test, divider))
    }

    let counts: number[] = []
    for (let i = 0; i < monkeys.length; i++) {
        counts.push(0)
    }

    for (let round = 0; round < 10000; round++) {
        for (let i = 0; i < monkeys.length; i++) {
            for (let item = 0; item < monkeys[i].items.length; item++) {
                counts[i]++
                let worry = monkeys[i].operation(monkeys[i].items[item])
                let nextMonkey = monkeys[i].test(worry)
                monkeys[nextMonkey].items.push(worry)
            }
            monkeys[i].items = []
        }
    }

    counts = counts.sort((a, b) => b - a)

    return counts[0] * counts[1]
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