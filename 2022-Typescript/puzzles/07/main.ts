import { readFile } from '../shared'
import fs from 'fs';

const day = "07"

class Directory {
    weight: number
    parent: string
    children: string[]

    constructor(parent: string) {
        this.weight = 0
        this.parent = parent
        this.children = []
    }
}

interface Dictionary<T> {
    [key: string]: T
}

function SolvePart1(inputFile: string) {
    let lines = readFile(inputFile).split("\n")

    let dirs: Dictionary<Directory> = {}

    let currentPath = ""
    for (let i = 0; i < lines.length; i++) {
        if (lines[i].split(" ")[1] == "cd") {
            if (lines[i].split(" ")[2] != "..") {
                dirs[currentPath + "." + lines[i].split(" ")[2]] = new Directory(currentPath)
                currentPath = currentPath + "." + lines[i].split(" ")[2]
            } else {
                currentPath = dirs[currentPath].parent
            }
        } else {
            if (lines[i].split(" ")[0] != "$") {
                if (lines[i].split(" ")[0] != "dir") {
                    dirs[currentPath].weight += +lines[i].split(" ")[0]
                    let currentParent = dirs[currentPath].parent
                    while (currentParent != "") {
                        dirs[currentParent].weight += +lines[i].split(" ")[0]
                        currentParent = dirs[currentParent].parent
                    }
                }
            }
        }
    }

    let total = 0
    for (let dir in dirs) {
        if (dirs[dir].weight <= 100000) {
            total += dirs[dir].weight
        }
    }

    return total
}

function SolvePart2(inputFile: string) {
    let lines = readFile(inputFile).split("\n")

    let dirs: Dictionary<Directory> = {}

    let currentPath = ""
    for (let i = 0; i < lines.length; i++) {
        if (lines[i].split(" ")[1] == "cd") {
            if (lines[i].split(" ")[2] != "..") {
                dirs[currentPath + "." + lines[i].split(" ")[2]] = new Directory(currentPath)
                currentPath = currentPath + "." + lines[i].split(" ")[2]
            } else {
                currentPath = dirs[currentPath].parent
            }
        } else {
            if (lines[i].split(" ")[0] != "$") {
                if (lines[i].split(" ")[0] != "dir") {
                    dirs[currentPath].weight += +lines[i].split(" ")[0]
                    let currentParent = dirs[currentPath].parent
                    while (currentParent != "") {
                        dirs[currentParent].weight += +lines[i].split(" ")[0]
                        currentParent = dirs[currentParent].parent
                    }
                }
            }
        }
    }

    let spaceToFree = 30000000 - (70000000 - dirs["./"].weight)
    let smallest = dirs["./"].weight

    for (let dir in dirs) {
        if (dirs[dir].weight >= spaceToFree) {
            if (dirs[dir].weight < smallest) {
                smallest = dirs[dir].weight
            }
        }
    }

    return smallest
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