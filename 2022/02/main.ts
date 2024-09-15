import { readFile } from '../shared'

const day = "02"
const fileContent = readFile(day + "/ab.input").split("\n")

let score = 0
for (let i = 0; i < fileContent.length; i++) {
    let line = fileContent[i]
    switch (line[2]) {
        case 'X':
            score += 1
            switch (line[0]) {
                case 'C':
                    score += 6
                    break
                case 'A':
                    score += 3
                    break
                default:
                    break
            }
            break
        case 'Y':
            score += 2
            switch (line[0]) {
                case 'A':
                    score += 6
                    break
                case 'B':
                    score += 3
                    break
                default:
                    break
            }
            break
        case 'Z':
            score += 3
            switch (line[0]) {
                case 'B':
                    score += 6
                    break
                case 'C':
                    score += 3
                    break
                default:
                    break
            }
            break
        default:
            break
    }
}

console.log("Part 1: " + score)

score = 0
for (let i = 0; i < fileContent.length; i++) {
    let line = fileContent[i]
    switch (line[2]) {
        case 'X':
            score += 0
            switch (line[0]) {
                case 'A':
                    score += 3
                    break
                case 'B':
                    score += 1
                    break
                default:
                    score += 2
                    break
            }
            break
        case 'Y':
            score += 3
            switch (line[0]) {
                case 'A':
                    score += 1
                    break
                case 'B':
                    score += 2
                    break
                default:
                    score += 3
                    break
            }
            break
        case 'Z':
            score += 6
            switch (line[0]) {
                case 'A':
                    score += 2
                    break
                case 'B':
                    score += 3
                    break
                default:
                    score += 1
                    break
            }
            break
        default:
            break
    }
}

console.log("Part 2: " + score)