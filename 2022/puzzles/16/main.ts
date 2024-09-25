import { readFile } from '../shared'
import fs from 'fs';

const day = "16"

function RecursiveWalk(currentValve: string, currentValue: number, time: number): number {
    console.log(currentValve)
    if (time < 0) {
        return currentValue
    }

    let max = 0
    dists.get(currentValve)!.forEach((dist, name) => {
        let nextValue = RecursiveWalk(name, currentValue + 1, time - dist)
        if (nextValue > max) {
            max = currentValue
        }
    })
    return max
}

let valveNames: string[] = []
let valves = new Map<string, number>()
let tunnels = new Map<string, string[]>()
let dists = new Map<string, Map<string, number>>()

function SolvePart1(inputFile: string) {
    let lines = readFile(inputFile).split("\n")

    for (let line of lines) {
        let name = line.split(" ")[1]
        let flow = +line.split("=")[1].split(";")[0]
        let tun = line.split("lead")[1].substring(" to valves ".length).split(", ")

        valveNames.push(name)
        valves.set(name, flow)
        tunnels.set(name, tun)
    }

    for (let valve of valveNames) {
        if (valve != "AA" && valves.get(valve) != 0) {
            continue
        }

        dists.set(valve, new Map<string, number>())
        dists.set(valve, dists.get(valve)!.set(valve, 0))
        dists.set(valve, dists.get(valve)!.set("AA", 0))
        let visited: string[] = [valve]

        let queue: [number, string][] = [[0, valve]]

        while (queue.length != 0) {
            let item = queue.pop()
            for (let neighbor of tunnels.get(item![1])!) {
                if (visited.includes(neighbor)) {
                    continue
                }
                visited.push(neighbor)
                if (valves.get(neighbor) != 0) {
                    dists.set(valve, dists.get(valve)!.set(neighbor, item![0] + 1))
                }
                queue.push([item![0] + 1, neighbor])
                queue.sort((a, b) => a[0] - b[0])
            }
        }

        dists.get(valve)!.delete(valve)
        if (valve != "AA") {
            dists.get(valve)!.delete("AA")
        }
    }

    console.log(dists)

    return RecursiveWalk("AA", 0, 30)
}

function SolvePart2(inputFile: string) {

    return 0
}

console.log("== TEST ==")
console.log(" -Part 1: " + SolvePart1(day + "/a.example"))
if (fs.existsSync(day + "/b.example")) {
    console.log(" -Part 2: " + SolvePart2(day + "/b.example"))
} else {
    console.log(" -Part 2: " + SolvePart2(day + "/a.example"))
}

if (false) {
    console.log("== DATA ==")
    console.log(" -Part 1: " + SolvePart1(day + "/ab.input"))
    console.log(" -Part 2: " + SolvePart2(day + "/ab.input"))
}