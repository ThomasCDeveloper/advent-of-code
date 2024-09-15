import { readFile } from '../shared'

const day = "01"
const fileContent = readFile(day + "/ab.input").split("\n")

let values: number[] = [0]

for (let i = 0; i < fileContent.length; i++) {
    const element = fileContent[i]

    if (element == "") {
        values.push(0)
    } else {
        values[values.length - 1] += +element
    }
}

values = values.sort((a, b) => b - a)

console.log("Part 1: " + values[0])
console.log("Part 2: " + (values[0] + values[1] + values[2]))
