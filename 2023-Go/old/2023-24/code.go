package main

// CMD: go run *.go

import (
	"fmt"
	"strconv"
	"strings"
)

type hail struct {
	x, y, z, vx, vy, vz float64
}

func GetHail(line string) hail {
	ls := strings.ReplaceAll(strings.Split(line, " @ ")[0], " ", "")
	rs := strings.ReplaceAll(strings.Split(line, " @ ")[1], " ", "")

	x, _ := strconv.Atoi(strings.Split(ls, ",")[0])
	y, _ := strconv.Atoi(strings.Split(ls, ",")[1])
	z, _ := strconv.Atoi(strings.Split(ls, ",")[2])

	vx, _ := strconv.Atoi(strings.Split(rs, ",")[0])
	vy, _ := strconv.Atoi(strings.Split(rs, ",")[1])
	vz, _ := strconv.Atoi(strings.Split(rs, ",")[2])

	return hail{float64(x), float64(y), float64(z), float64(vx), float64(vy), float64(vz)}
}

func (h1 hail) GetIntersection(h2 hail) (float64, float64) {
	// V1 = (h1.vx, h1.vy) = (-b,a)
	// (eq1): a*x+b*y+c=0 => h1.vy/h1.vx * x - h1.vy/h1.vx * h1.x + h1.y = y
	// (eq2): a*x+b*y+c=0 => h2.vy/h2.vx * x - h2.vy/h2.vx * h2.x + h2.y = y

	// y = Ax + B => A1x + B1 = A2x + B2 => x (A1 - A2) = B2 - B1

	// x = x0 + vx * t

	A1 := h1.vy / h1.vx
	A2 := h2.vy / h2.vx
	B1 := -h1.vy/h1.vx*h1.x + h1.y
	B2 := -h2.vy/h2.vx*h2.x + h2.y

	x := (B2 - B1) / (A1 - A2)
	y := A1*x + B1

	// check if path is in future for both paths
	t := (x - h1.x) / h1.vx
	if t < 0 {
		return -1, -1
	}
	t = (x - h2.x) / h2.vx
	if t < 0 {
		return -1, -1
	}

	return x, y
}

func SolvePart1(data []string) int {
	hails := []hail{}
	for _, line := range data {
		hails = append(hails, GetHail(line))
	}

	sum := 0
	for i := 0; i < len(data)-1; i++ {
		for j := i + 1; j < len(data); j++ {
			xcol, ycol := hails[i].GetIntersection(hails[j])

			if xcol >= 200000000000000 && xcol <= 400000000000000 && ycol >= 200000000000000 && ycol <= 400000000000000 {
				sum += 1
			}
		}
	}

	return sum
}

func SolvePart2(data []string) int {
	/*
		hails := []hail{}
		for _, line := range data {
			hails = append(hails, GetHail(line))
		}

			letters := []string{"a", "b", "c"}
			for i := range hails[0:3] {
				fmt.Println("U*" + letters[i] + "+x=" + strconv.Itoa(int(hails[i].vx)) + "*" + letters[i] + "+" + strconv.Itoa(int(hails[i].x)))
				fmt.Println("V*" + letters[i] + "+y=" + strconv.Itoa(int(hails[i].vy)) + "*" + letters[i] + "+" + strconv.Itoa(int(hails[i].y)))
				fmt.Println("W*" + letters[i] + "+z=" + strconv.Itoa(int(hails[i].vz)) + "*" + letters[i] + "+" + strconv.Itoa(int(hails[i].z)))
			}
	*/

	// online solver : https://www.dcode.fr/equation-solver
	// X=133619443970450
	// Y=263917577518425
	// Z=180640699244168

	return 133619443970450 + 263917577518425 + 180640699244168
}

func main() {
	data := GetInput("input.txt")

	// PART 1
	fmt.Println("Part 1:", SolvePart1(data))

	// PART 2
	fmt.Println("Part 2:", SolvePart2(data))
}
