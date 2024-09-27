import { Recoverable } from 'repl';
import { readFile } from '../shared'
import fs from 'fs';

const day = "16"

function SolvePart1(inputFile: string) {
    let lines = readFile(inputFile).split("\n")

    let valves: Record<string, number> = {}
    let tunnels: Record<string, string[]> = {}

    for (let line of lines) {
        let name = line.split(" ")[1]
        let flow = +line.split("=")[1].split(";")[0]
        let tun = line.split("lead")[1].substring(" to valves ".length).split(", ")

        valves[name] = flow
        tunnels[name] = tun
    }

    let dists: Record<string, Record<string, number>> = {}
    let nonEmpty: string[] = []

    for (let valve in valves) {
        if (valve != "AA" && valves[valve] == 0) continue
        if (valve != "AA") nonEmpty.push(valve)

        dists[valve] = { [valve]: 0, "AA": 0 }
        let visited = new Set<string>([valve])

        let queue: [number, string][] = []
        queue.push([0, valve])

        while (queue.length > 0) {
            let [dist, pos] = queue.shift()!
            for (let neighbor of tunnels[pos]) {
                if (visited.has(neighbor)) continue
                visited.add(neighbor)
                if (valves[neighbor]) {
                    dists[valve][neighbor] = dist + 1
                }
                queue.push([dist + 1, neighbor])
            }
        }

        delete dists[valve][valve]
        if (valve != "AA") delete dists[valve]["AA"]
    }

    let indices: Record<string, number> = {}

    nonEmpty.forEach((element, index) => {
        indices[element] = index
    })

    let cache: Record<string, number> = {}

    function dfs(timeLeft: number, valve: string, bitmask: number): number {
        let cacheKey = `${timeLeft}-${valve}-${bitmask}`
        if (cacheKey in cache) return cache[cacheKey]

        let maxVal = 0
        for (let neighbor in dists[valve]) {
            let bit = 1 << indices[neighbor]
            if (bitmask & bit) continue

            let nextTime = timeLeft - dists[valve][neighbor] - 1
            if (nextTime <= 0) continue

            maxVal = Math.max(maxVal, dfs(nextTime, neighbor, bitmask | bit) + valves[neighbor] * nextTime)
        }

        cache[cacheKey] = maxVal
        return maxVal
    }

    return dfs(30, "AA", 0)
}

function SolvePart2(inputFile: string) {
    let lines = readFile(inputFile).split("\n")

    let valves: Record<string, number> = {}
    let tunnels: Record<string, string[]> = {}

    for (let line of lines) {
        let name = line.split(" ")[1]
        let flow = +line.split("=")[1].split(";")[0]
        let tun = line.split("lead")[1].substring(" to valves ".length).split(", ")

        valves[name] = flow
        tunnels[name] = tun
    }

    let dists: Record<string, Record<string, number>> = {}
    let nonEmpty: string[] = []

    for (let valve in valves) {
        if (valve != "AA" && valves[valve] == 0) continue
        if (valve != "AA") nonEmpty.push(valve)

        dists[valve] = { [valve]: 0, "AA": 0 }
        let visited = new Set<string>([valve])

        let queue: [number, string][] = []
        queue.push([0, valve])

        while (queue.length > 0) {
            let [dist, pos] = queue.shift()!
            for (let neighbor of tunnels[pos]) {
                if (visited.has(neighbor)) continue
                visited.add(neighbor)
                if (valves[neighbor]) {
                    dists[valve][neighbor] = dist + 1
                }
                queue.push([dist + 1, neighbor])
            }
        }

        delete dists[valve][valve]
        if (valve != "AA") delete dists[valve]["AA"]
    }

    let indices: Record<string, number> = {}

    nonEmpty.forEach((element, index) => {
        indices[element] = index
    })

    let memo: Record<string, number> = {}

    function dfs(timeLeft: number, valve: string, bitmask: number): number {
        let cacheKey = `${timeLeft}-${valve}-${bitmask}`
        if (cacheKey in memo) return memo[cacheKey]

        let maxVal = 0
        for (let neighbor in dists[valve]) {
            let bit = 1 << indices[neighbor]
            if (bitmask & bit) continue

            let nextTime = timeLeft - dists[valve][neighbor] - 1
            if (nextTime <= 0) continue

            maxVal = Math.max(maxVal, dfs(nextTime, neighbor, bitmask | bit) + valves[neighbor] * nextTime)
        }

        memo[cacheKey] = maxVal
        return maxVal
    }

    let b = (1 << nonEmpty.length) - 1
    let max = 0
    for (let i = 0; i < b + 1; i++) {
        max = Math.max(dfs(26, "AA", i) + dfs(26, "AA", b ^ i), max)
    }

    return max
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