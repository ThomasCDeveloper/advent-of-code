import { readFile } from '../shared'
import fs from 'fs';

const day = "15"

function Manhattan(x0: number, y0: number, x1: number, y1: number): number {
    return Math.abs(x0 - x1) + Math.abs(y0 - y1)
}

function PushRange(r: [number, number], ranges: [number, number][]): [number, number][] {
    let intersecting: [number, number][] = []
    let notIntersecting: [number, number][] = []
    for (let range of ranges) {
        if ((r[0] >= range[0] && r[0] <= range[1] + 1) || (r[1] >= range[0] - 1 && r[0] <= range[1])) {
            intersecting.push(range)
        } else {
            notIntersecting.push(range)
        }
    }
    if (intersecting.length == 0) {
        ranges.push(r)
        return ranges.sort((a, b) => a[0] - b[0])
    }

    let min = r[0]
    let max = r[1]

    for (let range of intersecting) {
        if (range[0] < min) {
            min = range[0]
        }
        if (range[1] > max) {
            max = range[1]
        }
    }
    notIntersecting.push([min, max])

    return notIntersecting.sort((a, b) => a[0] - b[0])
}

function SolvePart1(inputFile: string, lineToCheck: number) {
    let lines = readFile(inputFile).split("\n").map((line) => [
        +line.split(",")[0].split("=")[1],
        +line.split(":")[0].split("=")[2],
        +line.split(",")[1].split("=")[2],
        +line.split(",")[2].split("=")[1]
    ])

    let sensors = lines.filter((xyxy) => Manhattan(xyxy[0], xyxy[1], xyxy[2], xyxy[3]) >= Math.abs(xyxy[1] - lineToCheck)).map((xyxy) => [
        Manhattan(xyxy[0], xyxy[1], xyxy[2], xyxy[3]) - Math.abs(xyxy[1] - lineToCheck) + 1,
        xyxy[0]
    ])

    let beaconsOnLine: number[] = []
    for (let line of readFile(inputFile).split("\n")) {
        if (+line.split(",")[2].split("=")[1] == lineToCheck) {
            beaconsOnLine.push(+line.split(",")[2].split("=")[1])
        }
    }

    let ranges: [number, number][] = []

    for (let sensor of sensors) {
        ranges = PushRange([sensor[1] - sensor[0] + 1, sensor[1] + sensor[0] - 1], ranges)
    }

    let total = 0
    for (let range of ranges) {
        total += range[1] - range[0]
    }

    return total - [... new Set(beaconsOnLine)].length - 1
}

function SolvePart2(inputFile: string, n: number) {
    let lines = readFile(inputFile).split("\n").map((line) => [
        +line.split(",")[0].split("=")[1],
        +line.split(":")[0].split("=")[2],
        +line.split(",")[1].split("=")[2],
        +line.split(",")[2].split("=")[1]
    ])
    for (let i = 0; i < n; i++) {
        let sensors = lines.filter((xyxy) => Manhattan(xyxy[0], xyxy[1], xyxy[2], xyxy[3]) >= Math.abs(xyxy[1] - i)).map((xyxy) => [
            Manhattan(xyxy[0], xyxy[1], xyxy[2], xyxy[3]) - Math.abs(xyxy[1] - i) + 1,
            xyxy[0]
        ])

        let beaconsOnLine: number[] = []
        for (let line of readFile(inputFile).split("\n")) {
            if (+line.split(",")[2].split("=")[1] == i) {
                beaconsOnLine.push(+line.split(",")[2].split("=")[1])
            }
        }

        let ranges: [number, number][] = []

        for (let sensor of sensors) {
            ranges = PushRange([sensor[1] - sensor[0] + 1, sensor[1] + sensor[0] - 1], ranges)
        }
        if (ranges.length > 1) {
            return i + (ranges[0][1] + 1) * 4000000
        }
    }

    return 0
}

console.log("== TEST ==")
console.log(" -Part 1: " + SolvePart1(day + "/a.example", 10))
if (fs.existsSync(day + "/b.example")) {
    console.log(" -Part 2: " + SolvePart2(day + "/b.example", 20))
} else {
    console.log(" -Part 2: " + SolvePart2(day + "/a.example", 20))
}

if (true) {
    console.log("== DATA ==")
    console.log(" -Part 1: " + SolvePart1(day + "/ab.input", 2000000))
    console.log(" -Part 2: " + SolvePart2(day + "/ab.input", 4000000))
}