import { sign } from 'crypto';
import { readFile } from '../shared'
import fs from 'fs';

const day = "09"

function SolvePart1(inputFile: string) {
    let lines = readFile(inputFile).split("\n")

    let headx = 0
    let heady = 0
    let tailx = 0
    let taily = 0

    let visitedPos: string[] = []
    for (let l = 0; l < lines.length; l++) {
        for (let i = 0; i < +lines[l].split(" ")[1]; i++) {
            switch (lines[l].split(" ")[0]) {
                case "U":
                    heady++
                    break
                case "D":
                    heady--
                    break
                case "R":
                    headx++
                    break
                default:
                    headx--
                    break
            }

            let diffx = headx - tailx
            let diffy = heady - taily

            if (Math.abs(diffx) > 1) {
                tailx += Math.sign(diffx)
                if (diffy != 0) {
                    taily += Math.sign(diffy)
                }
            }
            if (Math.abs(diffy) > 1) {
                taily += Math.sign(diffy)
                if (diffx != 0) {
                    tailx += Math.sign(diffx)
                }
            }

            if (visitedPos.indexOf(tailx + ":" + taily) < 0) {
                visitedPos.push(tailx + ":" + taily)
            }
        }
    }

    return visitedPos.length
}

function SolvePart2(inputFile: string) {
    let lines = readFile(inputFile).split("\n")
    let pos: number[][] = []

    let n = 10

    for (let i = 0; i < n; i++) {
        pos.push([0, 0])
    }

    let visitedPos: string[] = []

    for (let l = 0; l < lines.length; l++) {
        for (let i = 0; i < +lines[l].split(" ")[1]; i++) {
            switch (lines[l].split(" ")[0]) {
                case "U":
                    pos[0][1]++
                    break
                case "D":
                    pos[0][1]--
                    break
                case "R":
                    pos[0][0]++
                    break
                default:
                    pos[0][0]--
                    break
            }

            for (let j = 1; j < n; j++) {
                let headx = pos[j - 1][0]
                let heady = pos[j - 1][1]
                let tailx = pos[j][0]
                let taily = pos[j][1]

                let diffx = headx - tailx
                let diffy = heady - taily

                if (Math.abs(diffx) > 1 && Math.abs(diffy) > 1) {
                    pos[j][0] += Math.sign(diffx)
                    pos[j][1] += Math.sign(diffy)
                } else if (Math.abs(diffx) > 1) {
                    pos[j][0] += Math.sign(diffx)
                    if (diffy != 0) {
                        pos[j][1] += Math.sign(diffy)
                    }
                } else if (Math.abs(diffy) > 1) {
                    pos[j][1] += Math.sign(diffy)
                    if (diffx != 0) {
                        pos[j][0] += Math.sign(diffx)
                    }
                }
            }

            if (visitedPos.indexOf(pos[n - 1][0] + ":" + pos[n - 1][1]) < 0) {
                visitedPos.push(pos[n - 1][0] + ":" + pos[n - 1][1])
            }
        }
    }

    return visitedPos.length
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