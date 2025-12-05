import { readFile } from '../shared'
import fs from 'fs';

const day = "22"

let map: string[][]
let directions: Record<number, number[]> = {
    0: [1, 0],
    1: [0, 1],
    2: [-1, 0],
    3: [0, -1],
}

function Move(x: number, y: number, rot: number): [number, number] {
    let dir = directions[rot]

    if (x + dir[0] >= 0 && x + dir[0] < map[y].length && y + dir[1] >= 0 && y + dir[1] < map.length) {
        if (map[y + dir[1]][x + dir[0]] == "#") return [x, y]
        if (map[y + dir[1]][x + dir[0]] == ".") return [x + dir[0], y + dir[1]]
    }

    if (rot == 0) {
        for (let i = 0; i < map[y].length; i++) {
            if (map[y][i] == "#") return [x, y]
            if (map[y][i] == ".") return [i, y]
        }
    }
    if (rot == 2) {
        for (let i = map[y].length - 1; i >= 0; i--) {
            if (map[y][i] == "#") return [x, y]
            if (map[y][i] == ".") return [i, y]
        }
    }
    if (rot == 1) {
        for (let i = 0; i < map.length; i++) {
            if (map[i][x] == "#") return [x, y]
            if (map[i][x] == ".") return [x, i]
        }
    }
    if (rot == 3) {
        for (let i = map.length - 1; i >= 0; i--) {
            if (map[i][x] == "#") return [x, y]
            if (map[i][x] == ".") return [x, i]
        }
    }

    return [x + dir[0], y + dir[1]]
}

function SolvePart1(inputFile: string) {
    let [m, pw] = readFile(inputFile).split("\n\n")

    let password: (number | string)[] = []
    let currentNumber = 0
    for (let c of pw) {
        if (c == "R" || c == "L") {
            if (currentNumber != 0) {
                password.push(currentNumber)
                currentNumber = 0
            }
            password.push(c)
        } else {
            currentNumber *= 10
            currentNumber += +c
        }
    }
    if (currentNumber != 0) {
        password.push(currentNumber)
    }

    map = m.split("\n").map((line) => line.split(""))

    let [x, y, rot] = [map[0].indexOf("."), 0, 0]

    for (let command of password) {
        if (command === "R") {
            rot = (rot + 1) % 4
        } else if (command === "L") {
            rot = (rot + 3) % 4
        } else {
            for (let i = 0; i < (command as number); i++) {
                [x, y] = Move(x, y, rot)
            }
        }
    }

    return 1000 * (y + 1) + 4 * (x + 1) + rot
}

function SolvePart2(inputFile: string) {
    let [m, pw] = readFile(inputFile).split("\n\n")

    let password: (number | string)[] = []
    let currentNumber = 0
    for (let c of pw) {
        if (c == "R" || c == "L") {
            if (currentNumber != 0) {
                password.push(currentNumber)
                currentNumber = 0
            }
            password.push(c)
        } else {
            currentNumber *= 10
            currentNumber += +c
        }
    }
    if (currentNumber != 0) {
        password.push(currentNumber)
    }

    map = m.split("\n").map((line) => line.split(""))

    let [x, y, rot] = [map[0].indexOf("."), 0, 0]
    for (let command of password) {
        if (command === "R") {
            rot = (rot + 1) % 4
        } else if (command === "L") {
            rot = (rot + 3) % 4
        } else {
            for (let i = 0; i < (command as number); i++) {
                let [oldX, oldY, oldRot] = [x, y, rot]
                let dir = directions[rot];
                [x, y, rot] = [x + dir[0], y + dir[1], rot]

                if (y < 0 && x >= 50 && x < 100 && rot == 3) [x, y, rot] = [0, x + 100, 0]
                else if (x < 0 && y >= 150 && y < 200 && rot == 2) [x, y, rot] = [y - 100, 0, 1]
                else if (y < 0 && x >= 100 && x < 150 && rot == 3) [x, y, rot] = [x - 100, 199, 3]
                else if (y >= 200 && x >= 0 && x < 50 && rot == 1) [x, y, rot] = [x + 100, 0, 1]
                else if (x >= 150 && y >= 0 && y < 50 && rot == 0) [x, y, rot] = [99, 149 - y, 2]
                else if (x == 100 && y >= 100 && y < 150 && rot == 0) [x, y, rot] = [149, 149 - y, 2]
                else if (y == 50 && x >= 100 && x < 150 && rot == 1) [x, y, rot] = [99, x - 50, 2]
                else if (x == 100 && y >= 50 && y < 100 && rot == 0) [x, y, rot] = [y + 50, 49, 3]
                else if (y == 150 && x >= 50 && x < 100 && rot == 1) [x, y, rot] = [49, x + 100, 2]
                else if (x == 50 && y >= 150 && y < 200 && rot == 0) [x, y, rot] = [y - 100, 149, 3]
                else if (y == 99 && x >= 0 && x < 50 && rot == 3) [x, y, rot] = [50, x + 50, 0]
                else if (x == 49 && y >= 50 && y < 100 && rot == 2) [x, y, rot] = [y - 50, 100, 1]
                else if (x == 49 && y >= 0 && y < 50 && rot == 2) [x, y, rot] = [0, 149 - y, 0]
                else if (x < 0 && y >= 100 && y < 150 && rot == 2) [x, y, rot] = [50, 149 - y, 0]

                if (map[y][x] == " ") {
                    console.log("ERROR", x, y, rot)
                    return
                }

                if (map[y][x] == "#") {
                    [x, y, rot] = [oldX, oldY, oldRot]
                    break
                }
            }
        }
    }

    return 1000 * (y + 1) + 4 * (x + 1) + rot
}

if (true) {
    console.log("== DATA ==")
    console.log(" -Part 1: " + SolvePart1(day + "/ab.input"))
    console.log(" -Part 2: " + SolvePart2(day + "/ab.input"))
}