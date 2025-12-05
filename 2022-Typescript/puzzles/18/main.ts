import { NumericLiteral } from 'typescript';
import { readFile } from '../shared'
import fs from 'fs';

const day = "18"

function IsCubeInCubes(cube: number[], cubes: number[][]): boolean {
    for (let otherCube of cubes) {
        if (otherCube[0] == cube[0] && otherCube[1] == cube[1] && otherCube[2] == cube[2]) {
            return true
        }
    }
    return false
}

function SolvePart1(inputFile: string) {
    let cubes = readFile(inputFile).split("\n").map((l) => l.split(",").map((coord) => +coord))

    let faces = 0

    for (let cube of cubes) {
        for (let delta of [-1, 1]) {
            if (!IsCubeInCubes([cube[0] + delta, cube[1], cube[2]], cubes)) faces++
            if (!IsCubeInCubes([cube[0], cube[1] + delta, cube[2]], cubes)) faces++
            if (!IsCubeInCubes([cube[0], cube[1], cube[2] + delta], cubes)) faces++
        }
    }

    return faces
}

function IsAir(isCube: boolean[][][]): boolean[][][] {
    const directions = [
        [1, 0, 0], [-1, 0, 0],
        [0, 1, 0], [0, -1, 0],
        [0, 0, 1], [0, 0, -1]
    ]

    const inBounds = (x: number, y: number, z: number) =>
        x >= 0 && y >= 0 && z >= 0 &&
        x < isCube.length &&
        y < isCube[0].length &&
        z < isCube[0][0].length

    const reachable = Array(isCube.length)
        .fill(false)
        .map(() => Array(isCube[0].length)
            .fill(false)
            .map(() => Array(isCube[0][0].length).fill(false)))

    const visited = new Set<string>()
    const stack: [number, number, number][] = [[0, 0, 0]]

    while (stack.length > 0) {
        const [x, y, z] = stack.pop()!

        if (!inBounds(x, y, z) || isCube[x][y][z] || visited.has(`${x},${y},${z}`)) {
            continue
        }

        visited.add(`${x},${y},${z}`)
        reachable[x][y][z] = true

        for (const [dx, dy, dz] of directions) {
            stack.push([x + dx, y + dy, z + dz])
        }
    }

    return reachable
}

function SolvePart2(inputFile: string) {
    let cubes = readFile(inputFile).split("\n").map((l) => l.split(",").map((coord) => +coord))

    let maxX = Math.max(...cubes.map((cube) => cube[0])) + 1
    let minX = Math.min(...cubes.map((cube) => cube[0])) - 1

    let maxY = Math.max(...cubes.map((cube) => cube[1])) + 1
    let minY = Math.min(...cubes.map((cube) => cube[1])) - 1

    let maxZ = Math.max(...cubes.map((cube) => cube[2])) + 1
    let minZ = Math.min(...cubes.map((cube) => cube[2])) - 1

    let isCube: boolean[][][] = []
    for (let x = 0; x < maxX - minX; x++) {
        isCube.push([])
        for (let y = 0; y < maxY - minY; y++) {
            isCube[x].push([])
            for (let z = 0; z < maxZ - minZ; z++) {
                isCube[x][y].push(false)
            }
        }
    }
    for (let cube of cubes) {
        isCube[cube[0]][cube[1]][cube[2]] = true
    }

    let canReachAir = IsAir(isCube)
    let isInSolid: boolean[][][] = []
    let newCubes: number[][] = []
    for (let x = 0; x < maxX - minX; x++) {
        isInSolid.push([])
        for (let y = 0; y < maxY - minY; y++) {
            isInSolid[x].push([])
            for (let z = 0; z < maxZ - minZ; z++) {
                if (isCube[x][y][z] || !canReachAir[x][y][z]) {
                    newCubes.push([x, y, z])
                }
            }
        }
    }

    let faces = 0

    for (let cube of newCubes) {
        for (let delta of [-1, 1]) {
            if (!IsCubeInCubes([cube[0] + delta, cube[1], cube[2]], newCubes)) faces++
            if (!IsCubeInCubes([cube[0], cube[1] + delta, cube[2]], newCubes)) faces++
            if (!IsCubeInCubes([cube[0], cube[1], cube[2] + delta], newCubes)) faces++
        }
    }

    return faces
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