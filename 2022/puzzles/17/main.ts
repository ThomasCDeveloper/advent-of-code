import { readFile } from '../shared'
import fs from 'fs';

const day = "17"

function SolvePart1(inputFile: string) {
    let wind = readFile(inputFile).split("")

    let shapes: [number, number][][] = [
        [[0, 0], [1, 0], [2, 0], [3, 0]], // -
        [[1, 0], [0, 1], [1, 1], [2, 1], [1, 2]], // +
        [[0, 0], [1, 0], [2, 0], [2, 1], [2, 2]], // L
        [[0, 0], [0, 1], [0, 2], [0, 3]], // I
        [[0, 0], [1, 0], [0, 1], [1, 1]] // o
    ]

    let column: [number, number][] = []
    let maxHeight = 0

    let shapeIndex = 0
    let windIndex = 0

    while (shapeIndex < 2022) {
        let newRock = shapes[shapeIndex % 5].map((xy) => [xy[0] + 2, xy[1] + maxHeight + 3])

        let isLocked = false
        while (!isLocked) {
            // wind
            let offset = wind[windIndex % wind.length] == ">" ? 1 : -1
            let isBlocked = false
            for (let newRockPart of newRock) {
                if (newRockPart[0] + offset < 0 || newRockPart[0] + offset >= 7) {
                    isBlocked = true
                    break
                }
                for (let columnItem of column) {
                    if (columnItem[0] == newRockPart[0] + offset && columnItem[1] == newRockPart[1]) {
                        isBlocked = true
                    }
                }
            }

            if (!isBlocked)
                newRock = newRock.map((subRock) => [subRock[0] + offset, subRock[1]])

            windIndex++

            // gravity
            for (let newRockPart of newRock) {
                if (newRockPart[1] == 0) {
                    isLocked = true
                    break
                }
                for (let columnItem of column) {
                    if (columnItem[0] == newRockPart[0] && columnItem[1] == newRockPart[1] - 1) {
                        isLocked = true
                    }
                }
            }

            if (!isLocked) {
                newRock = newRock.map((subRock) => [subRock[0], subRock[1] - 1])
            }
        }

        newRock.forEach((newRockPart) => {
            maxHeight = Math.max(maxHeight, newRockPart[1] + 1)
            column.push([newRockPart[0], newRockPart[1]])
        })

        column.filter((xy) => {
            xy[1] >= maxHeight - 20
        })

        shapeIndex++
    }

    return maxHeight
}

function SolvePart2(inputFile: string) {
    let wind = readFile(inputFile).split("")

    let shapes: [number, number][][] = [
        [[0, 0], [1, 0], [2, 0], [3, 0]], // -
        [[1, 0], [0, 1], [1, 1], [2, 1], [1, 2]], // +
        [[0, 0], [1, 0], [2, 0], [2, 1], [2, 2]], // L
        [[0, 0], [0, 1], [0, 2], [0, 3]], // I
        [[0, 0], [1, 0], [0, 1], [1, 1]] // o
    ]

    let column: [number, number][] = []
    let maxHeight = 0

    let shapeIndex = 0
    let windIndex = 0

    while (shapeIndex < 2022) {
        let newRock = shapes[shapeIndex % 5].map((xy) => [xy[0] + 2, xy[1] + maxHeight + 3])

        let isLocked = false
        while (!isLocked) {
            // wind
            let offset = wind[windIndex % wind.length] == ">" ? 1 : -1
            let isBlocked = false
            for (let newRockPart of newRock) {
                if (newRockPart[0] + offset < 0 || newRockPart[0] + offset >= 7) {
                    isBlocked = true
                    break
                }
                for (let columnItem of column) {
                    if (columnItem[0] == newRockPart[0] + offset && columnItem[1] == newRockPart[1]) {
                        isBlocked = true
                    }
                }
            }

            if (!isBlocked)
                newRock = newRock.map((subRock) => [subRock[0] + offset, subRock[1]])

            windIndex++

            // gravity
            for (let newRockPart of newRock) {
                if (newRockPart[1] == 0) {
                    isLocked = true
                    break
                }
                for (let columnItem of column) {
                    if (columnItem[0] == newRockPart[0] && columnItem[1] == newRockPart[1] - 1) {
                        isLocked = true
                    }
                }
            }

            if (!isLocked) {
                newRock = newRock.map((subRock) => [subRock[0], subRock[1] - 1])
            }
        }

        let o = maxHeight
        newRock.forEach((newRockPart) => {
            maxHeight = Math.max(maxHeight, newRockPart[1] + 1)
            column.push([newRockPart[0], newRockPart[1]])
        })

        console.log(maxHeight - o)
        column.filter((xy) => {
            xy[1] >= maxHeight - 20
        })
        shapeIndex++
    }

    return maxHeight
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