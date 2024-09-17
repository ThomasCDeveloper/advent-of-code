import { readFile } from '../shared'
import fs from 'fs';

const day = "12"

let tiles: number[][] = []
let costs: number[][] = []
let startX = -1
let startY = -1
let endX = -1
let endY = -1

function WalkTile(posX: number, posY: number, cost: number) {
    cost += 1
    if (costs[posY][posX] <= cost) {
        return
    } else {
        costs[posY][posX] = cost
        if (posX > 0) {
            if (tiles[posY][posX - 1] - tiles[posY][posX] <= 1) {
                WalkTile(posX - 1, posY, cost)
            }
        }
        if (posY > 0) {
            if (tiles[posY - 1][posX] - tiles[posY][posX] <= 1) {
                WalkTile(posX, posY - 1, cost)
            }
        }
        if (posX < tiles[0].length - 1) {
            if (tiles[posY][posX + 1] - tiles[posY][posX] <= 1) {
                WalkTile(posX + 1, posY, cost)
            }
        }
        if (posY < tiles.length - 1) {
            if (tiles[posY + 1][posX] - tiles[posY][posX] <= 1) {
                WalkTile(posX, posY + 1, cost)
            }
        }
    }
}

function UnwalkTile(posX: number, posY: number, cost: number) {
    cost += 1
    if (costs[posY][posX] <= cost) {
        return
    } else {
        costs[posY][posX] = cost
        if (posX > 0) {
            if (tiles[posY][posX - 1] - tiles[posY][posX] >= -1) {
                UnwalkTile(posX - 1, posY, cost)
            }
        }
        if (posY > 0) {
            if (tiles[posY - 1][posX] - tiles[posY][posX] >= -1) {
                UnwalkTile(posX, posY - 1, cost)
            }
        }
        if (posX < tiles[0].length - 1) {
            if (tiles[posY][posX + 1] - tiles[posY][posX] >= -1) {
                UnwalkTile(posX + 1, posY, cost)
            }
        }
        if (posY < tiles.length - 1) {
            if (tiles[posY + 1][posX] - tiles[posY][posX] >= -1) {
                UnwalkTile(posX, posY + 1, cost)
            }
        }
    }
}

function SolvePart1(inputFile: string) {
    let lines = readFile(inputFile).split("\n")
    let abc123 = "abcdefghijklmnopqrstuvwxyz"

    for (let l = 0; l < lines.length; l++) {
        tiles.push([])
        costs.push([])
        for (let c = 0; c < lines[l].length; c++) {
            costs[l].push(lines.length * lines[0].length + 1)
            let char = lines[l][c]
            if (char == "S") {
                tiles[l].push(0)
                startX = c
                startY = l
            } else if (char == "E") {
                tiles[l].push(25)
                endX = c
                endY = l
            } else {
                tiles[l].push(abc123.indexOf(char))
            }
        }
    }

    let walkerX = startX
    let walkerY = startY

    WalkTile(walkerX, walkerY, -1)

    return costs[endY][endX]
}

function SolvePart2(inputFile: string) {
    let lines = readFile(inputFile).split("\n")
    costs = []
    for (let l = 0; l < tiles.length; l++) {
        costs.push([])
        for (let c = 0; c < tiles[l].length; c++) {
            costs[l].push(tiles.length * tiles[l].length + 1)
            let char = lines[l][c]
            if (char == "E") {
                startX = c
                startY = l
            }
        }
    }

    UnwalkTile(startX, startY, -1)

    let acosts: number[] = []

    for (let l = 0; l < tiles.length; l++) {
        for (let c = 0; c < tiles[l].length; c++) {
            if (tiles[l][c] == 0) {
                acosts.push(costs[l][c])
            }
        }
    }

    return Math.min(...acosts)
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