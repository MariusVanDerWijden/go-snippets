package terragen

import (
	"fmt"
	"math/rand"
)

const (
	EARTH = 77
	AIR   = 32
	WATER = 79
)

func addWater(field [][]byte, x, y int) [][]byte {
	if x < 0 || y < 0 || x > len(field) || y > len(field[0]) {
		return field
	}
	if field[x][y] == AIR {
		field[x][y] = WATER
		if field[x+1][y] == EARTH {
			if field[x][y-1] == AIR {
				field = addWater(field, x, y-1)
			}
			if field[x][y+1] == AIR {
				return addWater(field, x, y+1)
			}
			return field
		}
		return addWater(field, x+1, y)
	}
	if field[x][y] == WATER {
		return field
	}
	if field[x][y] == EARTH {
		if field[x][y-1] == AIR {
			return addWater(field, x, y-1)
		}
		if field[x][y+1] == AIR {
			return addWater(field, x, y+1)
		}
	}
	return field
}

func makeCave(field [][]byte, x, y, size int) [][]byte {
	if x < 0 || y < 0 || x > len(field) || y > len(field[0]) {
		return field
	}
	field[x][y] = AIR
	if size == 0 {
		return field
	}
	if rand.Intn(2) == 0 {
		field = makeCave(field, x, y, size-1)
	}
	x = x - 1 + rand.Intn(2)
	y = y - 1 + rand.Intn(2)
	return makeCave(field, x, y, size-1)
}

func createMap(height, width int, field byte) [][]byte {
	res := make([][]byte, height)
	for i := range res {
		res[i] = make([]byte, width)
		for k := range res[i] {
			res[i][k] = field
		}
	}
	return res
}

func printMap(field [][]byte) {
	for i := range field {
		for k := range field[i] {
			fmt.Printf("%v", string(field[i][k]))
		}
		fmt.Println()
	}
}
