import { readFile } from '../shared'
import fs from 'fs';

const day = "16"

// create a "state" class with time remaining, current pressure, map of the tunnels (without broken valves)
// when a valve is opened, break it and recalculate the map
// then, bruteforce it
const memo: Map<string, number> = new Map()

function FindShortestPath(start: string, end: string, tunnels: Map<string, string[]>): number {

    function dfs(node: string, target: string, visited: Set<string>): number {
        const key = `${node}-${target}`
        if (node === target) return 0
        if (memo.has(key)) return memo.get(key)!
        if (visited.has(node)) return Infinity
        visited.add(node)

        const neighbors = tunnels.get(node)
        if (!neighbors || neighbors.length === 0) {
            visited.delete(node)
            return Infinity
        }

        let minDistance = Infinity

        for (const neighbor of neighbors) {
            const distance = dfs(neighbor, target, visited)
            if (distance !== Infinity) {
                minDistance = Math.min(minDistance, distance + 1)
            }
        }
        visited.delete(node)
        memo.set(key, minDistance)

        return minDistance
    }

    const result = dfs(start, end, new Set<string>());
    return result === Infinity ? -1 : result;
}

function CalculateDists(currentvalve: string = "", valves: Map<string, number>) {
    let dists = new Map<string, Map<string, number>>()
    for (let valve1 of valveNames) {
        if (valves.get(valve1) == 0 && valve1 != currentvalve) continue
        if (dists.get(valve1) == undefined) {
            dists.set(valve1, new Map<string, number>())
        }
        for (let valve2 of valveNames) {
            if (valve1 === valve2) continue
            if (valves.get(valve2) == 0) continue

            if (dists.get(valve2) == undefined) {
                dists.set(valve2, new Map<string, number>())
            }

            let dist = FindShortestPath(valve1, valve2, tunnels)

            dists.set(valve1, dists.get(valve1)!.set(valve2, dist))
        }
    }
    return dists
}

class State {
    timeLeft: number
    pressure: number
    currentValve: string
    valves: Map<string, number>
    dists: Map<string, Map<string, number>>

    constructor(currentValve: string, timeLeft: number, pressure: number, valves: Map<string, number>, dists: Map<string, Map<string, number>>) {
        this.timeLeft = timeLeft
        this.pressure = pressure
        this.currentValve = currentValve
        this.valves = valves
        this.dists = dists
    }

    GetMaxPressure(): number {
        if (this.timeLeft <= 0) return this.pressure

        let nextValves = (new Map(this.valves)).set(this.currentValve, 0)
        let nextDists = CalculateDists(this.currentValve, this.valves)
        let maxPressure = new State(this.currentValve, this.timeLeft - 1, this.pressure + this.timeLeft * this.valves.get(this.currentValve)!, nextValves, nextDists).GetMaxPressure()

        for (let neighbor of this.dists.get(this.currentValve)!) {
            //console.log(this.timeLeft + ": " + this.currentValve + " => " + neighbor[0])
            let value = new State(neighbor[0], this.timeLeft - neighbor[1], this.pressure, this.valves, this.dists).GetMaxPressure()
            if (value > maxPressure) {
                maxPressure = value
            }
        }

        return maxPressure
    }
}

let valveNames: string[] = []
let tunnels = new Map<string, string[]>()

function SolvePart1(inputFile: string) {
    let lines = readFile(inputFile).split("\n")
    let valves = new Map<string, number>()

    for (let line of lines) {
        let name = line.split(" ")[1]
        let flow = +line.split("=")[1].split(";")[0]
        let tun = line.split("lead")[1].substring(" to valves ".length).split(", ")

        valveNames.push(name)
        valves.set(name, flow)
        tunnels.set(name, tun)
    }

    let state = new State("AA", 30, 0, valves, CalculateDists("AA", valves))

    return state.GetMaxPressure()
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