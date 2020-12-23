package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	path, _ := os.Getwd()
	file, err := os.Open(path + "/day20/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	state := "tile"
	var lineIndex int
	tiles := make(map[int64][][]byte)
	var currentID int64
	for scanner.Scan() {
		line := scanner.Text()
		switch state {
		case "tile":
			currentID, _ = strconv.ParseInt(strings.Replace(strings.Replace(line, "Tile ", "", 1), ":", "", 1), 10, 64)
			tiles[currentID] = [][]byte{}
			state = "line"
			lineIndex = 0
		case "line":
			tiles[currentID] = append(tiles[currentID], []byte(line))
			lineIndex++
			if lineIndex == 10 {
				state = ""
			}
		default:
			state = "tile"
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	part1(tiles)
}

func view(id int64, tile [][]byte) {
	fmt.Println(id)
	for _, row := range tile {
		for _, col := range row {
			fmt.Print(string(col))
		}
		fmt.Println()
	}
	fmt.Println()
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

func flipVert(tile [][]byte) [][]byte {
	result := [][]byte{}
	for _, line := range tile {
		result = append(result, reverse(line))
	}
	return result
}

func multipleRotate(tile [][]byte, times int) [][]byte {
	result := copyTile(tile)
	for i := 1; i <= times; i++ {
		result = rotate(result)
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

type coordinates struct {
	x, y int
}

func move(c coordinates, direction string) coordinates {
	newC := c
	switch direction {
	case "top":
		newC.y = c.y - 1
	case "bottom":
		newC.y = c.y + 1
	case "left":
		newC.x = c.y - 1
	case "right":
		newC.x = c.y + 1
	}
	return newC
}

func part1(tiles map[int64][][]byte) {
	picture := make(map[coordinates][][]byte)
	ids := make(map[coordinates]int64)
	coords := make(map[int64]coordinates)
	tilesToProcess := []int64{}
	for k, pic := range tiles {
		tilesToProcess = append(tilesToProcess, k)
		ids[coordinates{0, 0}] = k
		picture[coordinates{0, 0}] = pic
		coords[k] = coordinates{0, 0}
		break
	}
	for {
		fmt.Println(tilesToProcess)
		fmt.Println(ids)
		fmt.Println(coords)
		mainID := tilesToProcess[0]
		tilesToProcess = tilesToProcess[1:]
		_, exists := coords[mainID]
		if exists {

		}
		for id, tile := range tiles {
			if id != mainID {
				direction := ""
				for r := 0; r < 4; r++ {
					for f := 0; f < 4; f++ {
						result := multipleRotate(tile, r)
						result = multipleFlip(result, f)
						direction = match(tiles[mainID], result)
						if direction != "" {
							tileCoords := move(coords[mainID], direction)
							tilesToProcess = append(tilesToProcess, id)
							ids[tileCoords] = id
							picture[tileCoords] = result
							coords[id] = tileCoords
							break
						}
					}
					if direction != "" {
						break
					}
				}
				_, existsTop := ids[move(coords[mainID], "top")]
				_, existsBottom := ids[move(coords[mainID], "bottom")]
				_, existsLeft := ids[move(coords[mainID], "left")]
				_, existsRight := ids[move(coords[mainID], "right")]
				if existsBottom && existsTop && existsRight && existsLeft {
					break
				}
			}
		}
		if len(tilesToProcess) == 0 {
			break
		}
	}
	fmt.Println(ids)
}
