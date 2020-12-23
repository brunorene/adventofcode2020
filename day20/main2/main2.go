package main2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type sideMatch struct {
	position string
	rotation int
	flip     int
}

type tile struct {
	id       int64
	rotation int
	flip     int
	square   [][]byte
	sides    map[int64]sideMatch
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
	var current tile
	var lineIndex int
	tiles := make(map[int64]tile)
	for scanner.Scan() {
		line := scanner.Text()
		switch state {
		case "tile":
			id, _ := strconv.ParseInt(strings.Replace(strings.Replace(line, "Tile ", "", 1), ":", "", 1), 10, 64)
			current = tile{id: id, sides: make(map[int64]sideMatch), square: [][]byte{}}
			state = "line"
			lineIndex = 0
		case "line":
			current.square = append(current.square, []byte(line))
			lineIndex++
			if lineIndex == 10 {
				state = ""
			}
		default:
			tiles[current.id] = current
			state = "tile"
		}
	}
	tiles[current.id] = current

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	part2(part1(tiles))
}

func reverse(b []byte) []byte {
	newB := b
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		newB[i], newB[j] = b[j], b[i]
	}
	return newB
}

func flipVert(t tile) tile {
	result := tile{id: t.id, rotation: t.rotation, sides: make(map[int64]sideMatch)}
	for _, line := range t.square {
		result.square = append(result.square, reverse(line))
	}
	for key, match := range result.sides {
		switch match.position {
		case "left":
			result.sides[key] = sideMatch{"right", match.rotation, match.flip}
		case "right":
			result.sides[key] = sideMatch{"left", match.rotation, match.flip}
		}
	}
	return result
}

func flipHorz(t tile) tile {
	result := rotate(t)
	result = flipVert(result)
	return multipleRotate(result, 3)
}

func multipleFlip(t tile, times int) tile {
	times = times % 4
	result := copyTile(t)
	switch times {
	case 1:
		result = flipVert(result)
	case 2:
		result = flipHorz(result)
	case 3:
		result = flipVert(result)
		result = flipHorz(result)
	}
	result.flip = times
	return result
}

func rotate(t tile) tile {
	result := tile{id: t.id, rotation: (t.rotation + 1) % 4, sides: make(map[int64]sideMatch)}
	for range t.square {
		result.square = append(result.square, make([]byte, len(t.square)))
	}
	n := len(t.square)
	x := n / 2
	y := n - 1
	for i := 0; i < x; i++ {
		for j := i; j < y-i; j++ {
			result.square[i][j] = t.square[y-j][i]
			result.square[y-j][i] = t.square[y-i][y-j]
			result.square[y-i][y-j] = t.square[j][y-i]
			result.square[j][y-i] = t.square[i][j]
		}
	}
	for key, match := range result.sides {
		switch match.position {
		case "bottom":
			t.sides[key] = sideMatch{"right", match.rotation, match.flip}
		case "top":
			t.sides[key] = sideMatch{"left", match.rotation, match.flip}
		case "left":
			t.sides[key] = sideMatch{"bottom", match.rotation, match.flip}
		case "right":
			t.sides[key] = sideMatch{"top", match.rotation, match.flip}
		}
	}
	return result
}

func multipleRotate(t tile, times int) tile {
	result := copyTile(t)
	if times > 0 {
		result = rotate(t)
	}
	for i := 1; i < times; i++ {
		result = rotate(result)
	}
	return result
}

func match(center tile, side tile) string {
	directions := make(map[string]int)
	for pos := range center.square {
		if center.square[0][pos] == side.square[9][pos] {
			directions["top"]++
		}
		if center.square[9][pos] == side.square[0][pos] {
			directions["bottom"]++
		}
		if center.square[pos][0] == side.square[pos][9] {
			directions["left"]++
		}
		if center.square[pos][9] == side.square[pos][0] {
			directions["right"]++
		}
	}
	for direction, count := range directions {
		if count == len(center.square) {
			return direction
		}
	}
	return ""
}

func view(t tile) {
	fmt.Println(t.id, t.rotation, t.flip, t.sides)
	for _, row := range t.square {
		for _, col := range row {
			fmt.Print(string(col))
		}
		fmt.Println()
	}
	fmt.Println()
}

func findSides(t tile, other map[int64]tile) tile {
	for _, side := range other {
		for r := 0; r < 4; r++ {
			for f := 0; f < 4; f++ {
				changed := multipleRotate(side, r)
				changed = multipleFlip(changed, f)
				position := match(t, changed)
				if position != "" {
					t.sides[changed.id] = sideMatch{position, changed.rotation, changed.flip}
					break
				}
			}
		}
	}
	return t
}

func copyMap(src map[int64]tile) map[int64]tile {
	dst := make(map[int64]tile)
	for k, v := range src {
		dst[k] = v
	}
	return dst
}

func copyTile(t tile) tile {
	result := tile{t.id, t.rotation, t.flip, [][]byte{}, make(map[int64]sideMatch)}
	for _, row := range t.square {
		copyRow := []byte{}
		for _, item := range row {
			copyRow = append(copyRow, item)
		}
		result.square = append(result.square, copyRow)
	}
	for k, v := range t.sides {
		result.sides[k] = v
	}
	return result
}

func remove(m map[int64]tile, key int64) map[int64]tile {
	result := copyMap(m)
	delete(result, key)
	return result
}

func part1(tiles map[int64]tile) map[int64]tile {
	mul := int64(1)
	for _, main := range tiles {
		other := remove(tiles, main.id)
		main = findSides(main, other)
		if len(main.sides) == 2 {
			mul *= main.id
		}
	}
	fmt.Println(mul)
	return tiles
}

func removeBorder(t tile) [][]byte {
	result := [][]byte{}
	for i := 1; i < len(t.square)-1; i++ {
		result = append(result, t.square[i][1:len(t.square)-1])
	}
	return result
}

func nextTile(t tile, position string, tiles map[int64]tile) *tile {
	for id, m := range t.sides {
		if m.position == position {
			result := tiles[id]
			return &result
		}
	}
	return nil
}

func draw(ascii [][]byte) {
	for _, line := range ascii {
		fmt.Println(line)
	}
}

func part2(tiles map[int64]tile) {
	var topLeft tile
	for _, t := range tiles {
		if len(t.sides) == 2 {
			sides := make(map[string]int)
			for _, m := range t.sides {
				sides[m.position] = 1
			}
			if sides["right"] == 1 && sides["bottom"] == 1 {
				topLeft = t
				break
			}
		}
	}

	completePicture := [][]byte{}

	current := &topLeft
	// going down
	for current != nil {
		fragment := removeBorder(*current)
		for _, line := range fragment {
			completePicture = append(completePicture, line)
		}
		old := current
		current = nextTile(*current, "bottom", tiles)
		if current != nil {
			next := nextTile(*current, "bottom", tiles)
			if next.id == current.id {
				tile := multipleFlip(*current, 1)
				current = &tile

			}
		}
		fmt.Println(old.id, current.id)
	}

	draw(completePicture)
}
