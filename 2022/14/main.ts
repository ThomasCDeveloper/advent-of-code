import internal from 'stream';
import { readFile } from '../shared'
import fs from 'fs';
import { off } from 'process';

const day = "14"

function PrintGrid(grid: boolean[][]) {
    for (let l of grid) {
        console.log(l.map((b) => b ? "#" : ".").join(""))
    }
}

function MoveGrain(grainX: number, grainY: number, grid: boolean[][]): [number, number, boolean[][], boolean] {
    if (grid[grainY + 1][grainX] == false) {
        return [grainX, grainY + 1, grid, false]
    }
    if (grid[grainY + 1][grainX - 1] == false) {
        return [grainX - 1, grainY + 1, grid, false]
    }
    if (grid[grainY + 1][grainX + 1] == false) {
        return [grainX + 1, grainY + 1, grid, false]
    }

    grid[grainY][grainX] = true
    return [0, 0, grid, true]
}

function SolvePart1(inputFile: string) {
    let paths = readFile(inputFile).split("\n").map((line) => line.split(" -> ").map((xy) => xy.split(",").map((xory) => +xory)))
    let maxX = 500
    let maxY = 0
    let minX = 500
    let minY = 0

    for (let path of paths) {
        for (let xy of path) {
            if (xy[0] > maxX) {
                maxX = xy[0]
            }
            if (xy[0] < minX) {
                minX = xy[0]
            }
            if (xy[1] > maxY) {
                maxY = xy[1]
            }
            if (xy[1] < minY) {
                minY = xy[1]
            }
        }
    }

    let offsetX = minX - 1
    let offsetY = minY - 1

    let grid: boolean[][] = []
    for (let y = 0; y < maxY - offsetY + 3; y++) {
        grid.push([])
        for (let x = 0; x < maxX - offsetX + 2; x++) {
            grid[y].push(false)
        }
    }

    for (let path of paths) {
        let walkerX = path[0][0]
        let walkerY = path[0][1]
        for (let i = 1; i < path.length; i++) {
            grid[walkerY - offsetY][walkerX - offsetX] = true
            while (walkerX != path[i][0] || walkerY != path[i][1]) {
                walkerX += Math.sign(path[i][0] - walkerX)
                walkerY += Math.sign(path[i][1] - walkerY)
                grid[walkerY - offsetY][walkerX - offsetX] = true
            }
        }
    }


    let grainX = 500 - offsetX
    let grainY = 0 - offsetY
    let grainIsResting = false
    let nbGrain = 0
    let a = 0
    while (grainY < maxY - offsetY) {
        if (grainIsResting) {
            grainX = 500 - offsetX
            grainY = 0 - offsetY
            nbGrain++
        }
        [grainX, grainY, grid, grainIsResting] = MoveGrain(grainX, grainY, grid)
    }

    return nbGrain
}

function SolvePart2(inputFile: string) {
    let paths = readFile(inputFile).split("\n").map((line) => line.split(" -> ").map((xy) => xy.split(",").map((xory) => +xory)))
    let maxX = 500
    let maxY = 0
    let minX = 500
    let minY = 0

    for (let path of paths) {
        for (let xy of path) {
            if (xy[0] > maxX) {
                maxX = xy[0]
            }
            if (xy[0] < minX) {
                minX = xy[0]
            }
            if (xy[1] > maxY) {
                maxY = xy[1]
            }
            if (xy[1] < minY) {
                minY = xy[1]
            }
        }
    }

    let offsetX = minX - 222
    let offsetY = minY - 1

    let grid: boolean[][] = []
    for (let y = 0; y < maxY - offsetY + 3; y++) {
        grid.push([])
        for (let x = 0; x < maxX - offsetX + 223; x++) {
            grid[y].push(y == maxY - offsetY + 2 ? true : false)
        }
    }

    for (let path of paths) {
        let walkerX = path[0][0]
        let walkerY = path[0][1]
        for (let i = 1; i < path.length; i++) {
            grid[walkerY - offsetY][walkerX - offsetX] = true
            while (walkerX != path[i][0] || walkerY != path[i][1]) {
                walkerX += Math.sign(path[i][0] - walkerX)
                walkerY += Math.sign(path[i][1] - walkerY)
                grid[walkerY - offsetY][walkerX - offsetX] = true
            }
        }
    }

    let lastNotStillX = 0
    let lastNotStillY = 0
    let grainX = 500 - offsetX
    let grainY = 0 - offsetY
    let grainIsResting = false
    let nbGrain = 0
    while (grid[0 - offsetY][500 - offsetX] != true) {
        if (grainIsResting) {
            grainX = 500 - offsetX
            grainY = 0 - offsetY
            nbGrain++
        }
        [grainX, grainY, grid, grainIsResting] = MoveGrain(grainX, grainY, grid)
        /* // Comment lines 150~155 and remove the colon on line 163
        if (grainIsResting) {
            grainX = lastNotStillX
            grainY = lastNotStillY
            nbGrain++
        }
        lastNotStillX = grainX
        lastNotStillY = grainY;
        [grainX, grainY, grid, grainIsResting] = MoveGrain(grainX, grainY, grid)
        */
    }

    PrintGrid(grid)

    return nbGrain + 1
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