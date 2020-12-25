package main

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type square struct {
	ID      int64
	x, y    int
	picture [][]byte
}

func main() {
	path, _ := os.Getwd()
	file, err := os.Open(path + "/day20/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	state := "tile"
	tiles := make(map[int64]square)
	var currentID int64
	for scanner.Scan() {
		line := scanner.Text()
		switch state {
		case "tile":
			currentID, _ = strconv.ParseInt(strings.Replace(strings.Replace(line, "Tile ", "", 1), ":", "", 1), 10, 64)
			tiles[currentID] = square{ID: currentID}
			state = "line"
		case "line":
			if line == "" {
				state = "tile"
				continue
			}
			square := tiles[currentID]
			square.picture = append(square.picture, []byte(line))
			tiles[currentID] = square
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	part2(part1(tiles))
}

func view(s square) {
	fmt.Println(s.ID, s.x, s.y)
	for _, row := range s.picture {
		fmt.Println(string(row))
	}
}

func rotate(tile [][]byte) [][]byte {
	result := [][]byte{}
	for range tile {
		result = append(result, make([]byte, len(tile)))
	}
	n := len(tile)
	x := n / 2
	y := n - 1
	for i := 0; i < x; i++ {
		for j := i; j < y-i; j++ {
			result[i][j] = tile[y-j][i]
			result[y-j][i] = tile[y-i][y-j]
			result[y-i][y-j] = tile[j][y-i]
			result[j][y-i] = tile[i][j]
		}
	}
	return result
}

func copyTile(tile [][]byte) [][]byte {
	result := [][]byte{}
	for _, row := range tile {
		copyRow := []byte{}
		for _, item := range row {
			copyRow = append(copyRow, item)
		}
		result = append(result, copyRow)
	}
	return result
}

func reverse(b []byte) []byte {
	newB := b
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		newB[i], newB[j] = b[j], b[i]
	}
	return newB
}

func multipleRotate(tile [][]byte, times int) [][]byte {
	result := copyTile(tile)
	for i := 1; i <= times; i++ {
		result = rotate(result)
	}
	return result
}

func flipVert(tile [][]byte) [][]byte {
	result := [][]byte{}
	for _, line := range tile {
		result = append(result, reverse(line))
	}
	return result
}

func flipHorz(tile [][]byte) [][]byte {
	result := rotate(tile)
	result = flipVert(result)
	return multipleRotate(result, 3)
}

func multipleFlip(tile [][]byte, times int) [][]byte {
	times = times % 4
	result := copyTile(tile)
	switch times {
	case 1:
		result = flipVert(result)
	case 2:
		result = flipHorz(result)
	case 3:
		result = flipVert(result)
		result = flipHorz(result)
	}
	return result
}

func match(center [][]byte, side [][]byte) string {
	directions := make(map[string]int)
	for pos := range center {
		if center[0][pos] == side[9][pos] {
			directions["top"]++
		}
		if center[9][pos] == side[0][pos] {
			directions["bottom"]++
		}
		if center[pos][0] == side[pos][9] {
			directions["left"]++
		}
		if center[pos][9] == side[pos][0] {
			directions["right"]++
		}
	}
	for direction, count := range directions {
		if count == len(center) {
			return direction
		}
	}
	return ""
}

func move(x, y int, direction string) (int, int) {
	switch direction {
	case "top":
		return x, y - 1
	case "bottom":
		return x, y + 1
	case "left":
		return x - 1, y
	case "right":
		return x + 1, y
	}
	return x, y
}

type transformation struct {
	rotate, flip int
}

func transformations() []transformation {
	result := []transformation{}
	for r := 0; r < 4; r++ {
		for f := 0; f < 4; f++ {
			result = append(result, transformation{r, f})
		}
	}
	return result
}

func addToGrid(grid map[int]map[int]square, x, y int, val square) map[int]map[int]square {
	_, exists := grid[x]
	if !exists {
		grid[x] = make(map[int]square)
	}
	grid[x][y] = val
	return grid
}

func part1(tiles map[int64]square) [][]square {
	processed := make(map[int64]bool)
	processQueue := list.New()
	max := int64(0)
	for ID := range tiles {
		if ID > max {
			max = ID
		}
	}
	currentSquare := tiles[max]
	currentSquare.x = 0
	currentSquare.y = 0
	tiles[max] = currentSquare
	for {
		nextSquares := make(map[string]int64)
		for ID, tile := range tiles {
			if ID != currentSquare.ID && !processed[ID] {
				for _, t := range transformations() {
					result := multipleFlip(multipleRotate(tile.picture, t.rotate), t.flip)
					placement := match(currentSquare.picture, result)
					if placement != "" {
						tile.x, tile.y = move(currentSquare.x, currentSquare.y, placement)
						tile.picture = result
						tiles[ID] = tile
						nextSquares[placement] = ID
						break
					}
				}
			}
		}
		processed[currentSquare.ID] = true
		for _, ID := range nextSquares {
			if !processed[ID] {
				processQueue.PushBack(ID)
			}
		}
		for {
			if processQueue.Len() == 0 {
				break
			}
			nextID := processQueue.Remove(processQueue.Front()).(int64)
			if !processed[nextID] {
				currentSquare = tiles[nextID]
				break
			}
		}
		if processQueue.Len() == 0 {
			break
		}
	}
	minX := math.MaxInt64
	minY := math.MaxInt64
	for _, tile := range tiles {
		if tile.x < minX {
			minX = tile.x
		}
		if tile.y < minY {
			minY = tile.y
		}
	}
	for ID, tile := range tiles {
		tile.x -= minX
		tile.y -= minY
		tiles[ID] = tile
	}
	maxX := 0
	maxY := 0
	for _, tile := range tiles {
		if tile.x > maxX {
			maxX = tile.x
		}
		if tile.y > maxY {
			maxY = tile.y
		}
	}
	grid := [][]square{}
	for y := 0; y <= maxY; y++ {
		grid = append(grid, []square{})
		for x := 0; x <= maxX; x++ {
			grid[y] = append(grid[y], square{})
		}
	}
	for _, tile := range tiles {
		grid[tile.y][tile.x] = tile
	}
	fmt.Println(grid[0][0].ID * grid[0][len(grid)-1].ID * grid[len(grid)-1][0].ID * grid[len(grid)-1][len(grid)-1].ID)
	return grid
}

func removeBorder(picture [][]byte) [][]byte {
	borderLess := [][]byte{}
	for row := 1; row < len(picture)-1; row++ {
		borderLess = append(borderLess, []byte{})
		for col := 1; col < len(picture[row])-1; col++ {
			borderLess[row-1] = append(borderLess[row-1], picture[row][col])
		}
	}
	return borderLess
}

func viewPicture(picture [][]byte) {
	fmt.Println(len(picture), "x", len(picture[0]))
	for _, row := range picture {
		fmt.Println(string(row))
	}
}

func intersection(a []int, b []int) []int {
	is := make(map[int]bool)
	for _, v := range a {
		is[v] = true
	}
	overlap := []int{}
	for _, v := range b {
		if is[v] {
			overlap = append(overlap, v)
		}
	}
	return overlap
}

func countChar(matrix [][]byte, c byte) int {
	count := 0
	for y := 0; y < len(matrix); y++ {
		for x := 0; x < len(matrix); x++ {
			if matrix[y][x] == c {
				count++
			}
		}
	}
	return count
}

type location struct {
	x, y int
}

func part2(grid [][]square) {
	completePicture := [][]byte{}
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			borderLess := removeBorder(grid[y][x].picture)
			for row := 0; row < len(borderLess); row++ {
				if x == 0 {
					completePicture = append(completePicture, []byte{})
				}
				for col := 0; col < len(borderLess[row]); col++ {
					completePicture[row+y*len(borderLess)] = append(completePicture[row+y*len(borderLess)], borderLess[row][col])
				}
			}
		}
	}
	regSharp := regexp.MustCompile(`#`)
	for _, t := range transformations() {
		fmt.Println("transformation", t)
		result := multipleFlip(multipleRotate(completePicture, t.rotate), t.flip)
		var index0, index1, index2 []int
		monsterStart := []location{}
		for row := 0; row < len(result)-2; row++ {
			findAll := regSharp.FindAllIndex(result[row], -1)
			index0 = []int{}
			for _, find := range findAll {
				idx := find[0] - 18
				if idx >= 0 {
					index0 = append(index0, idx)
				}
			}
			findAll = regSharp.FindAllIndex(result[row+1], -1)
			index1 = []int{}
			places := []int{0, 5, 6, 11, 12, 17, 18, 19}
			for _, find := range findAll {
				match := true
				for _, place := range places {
					nextIdx := find[0] + place
					if nextIdx >= len(result[row+1]) {
						match = false
						break
					}
					if result[row+1][find[0]+place] != '#' {
						match = false
					}
				}
				if match {
					index1 = append(index1, find[0])
				}
			}
			findAll = regSharp.FindAllIndex(result[row+2], -1)
			index2 = []int{}
			places = []int{3, 6, 9, 12, 15}
			for _, find := range findAll {
				idx := find[0] - 1
				match := true
				for _, place := range places {
					nextIdx := find[0] + place
					if nextIdx >= len(result[row+1]) {
						match = false
						break
					}
					if result[row+2][nextIdx] != '#' {
						match = false
					}
				}
				if match && idx >= 0 {
					index2 = append(index2, idx)
				}
			}
			common := intersection(index0, intersection(index1, index2))
			if len(common) > 0 {
				for _, x := range common {
					monsterStart = append(monsterStart, location{x, row})
				}
			}
		}
		if len(monsterStart) > 0 {
			monsterSize := 15
			total := countChar(result, byte('#'))
			fmt.Println(total - len(monsterStart)*monsterSize)
		}
	}
}
