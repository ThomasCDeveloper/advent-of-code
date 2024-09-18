import { readFile } from '../shared'
import fs from 'fs';

const day = "13"

type item = number | item[]

function isItemANumber(it: item): boolean {
    return !Array.isArray(it)
}

function CompareItems(it0: item, it1: item): boolean | null {
    let i0 = (it0 as item[])
    let i1 = (it1 as item[])

    for (let i = 0; i < i0.length && i < i1.length; i++) {
        if (isItemANumber(i0[i]) && isItemANumber(i1[i])) {
            let n0 = i0[i] as number
            let n1 = i1[i] as number

            if (n0 < n1) {
                return true
            }
            if (n0 > n1) {
                return false
            }
            continue
        }

        if (!isItemANumber(i0[i]) && !isItemANumber(i1[i])) {
            let comparison = CompareItems(i0[i], i1[i])
            if (comparison === null) {
                continue
            }
            return comparison
        }

        if (!isItemANumber(i0[i]) && isItemANumber(i1[i])) {
            let comparison = CompareItems(i0[i], [i1[i]])
            if (comparison === null) {
                continue
            }
            return comparison
        }

        if (isItemANumber(i0[i]) && !isItemANumber(i1[i])) {
            let comparison = CompareItems([i0[i]], i1[i])
            if (comparison === null) {
                continue
            }
            return comparison
        }
    }

    if (i0.length < i1.length) {
        return true
    }
    if (i0.length > i1.length) {
        return false
    }

    return null
}

function Parse(str: string): item[] {
    str = str.substring(1, str.length - 1)

    if (str.length == 0) {
        return []
    }

    let subStrings: string[] = []
    let currentString = ""
    for (let c = 0; c < str.length; c++) {
        if (str[c] == ",") {
            if (currentString != "") {
                subStrings.push(currentString)
            }
            currentString = ""
            continue
        }
        if (str[c] == "[") {
            let bracketsOpened = 1
            while (bracketsOpened > 0) {
                currentString += str[c]
                c++
                if (str[c] == "]") {
                    bracketsOpened--
                }
                if (str[c] == "[") {
                    bracketsOpened++
                }
            }
            subStrings.push(currentString + "]")
            currentString = ""
            continue
        }
        currentString += str[c]
    }
    if (currentString != "") {
        subStrings.push(currentString)
    }


    let output: item = []

    for (let s = 0; s < subStrings.length; s++) {
        if (subStrings[s].includes("[")) {
            output.push(Parse(subStrings[s]))
        } else {
            output.push(+subStrings[s])
        }
    }

    return output
}

function SolvePart1(inputFile: string) {
    let lines = readFile(inputFile).split("\n\n")

    let items: item[] = []

    for (let i = 0; i < lines.length; i++) {
        let i1 = lines[i].split("\n")[0]
        items.push(Parse(i1))
        let i2 = lines[i].split("\n")[1]
        items.push(Parse(i2))
    }

    let sum = 0
    for (let i = 0; i < items.length; i += 2) {
        if (CompareItems(items[i], items[i + 1])) {
            sum += 1 + i / 2
        }
    }

    return sum
}

function SolvePart2(inputFile: string) {
    let lines = readFile(inputFile).split("\n\n")

    let items: item[] = [[[2]], [[6]]]

    for (let i = 0; i < lines.length; i++) {
        let i1 = lines[i].split("\n")[0]
        items.push(Parse(i1))
        let i2 = lines[i].split("\n")[1]
        items.push(Parse(i2))
    }

    items.sort((a, b) => (CompareItems(a, b) as boolean) ? -1 : 1)

    let index2 = -1
    let index6 = -1
    for (let i = 0; i < items.length; i++) {
        if (CompareItems(items[i], [[2]]) == null) {
            index2 = i + 1
        }
        if (CompareItems(items[i], [[6]]) == null) {
            index6 = i + 1
        }
    }

    return index2 * index6
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