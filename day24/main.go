package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

type piece struct {
	color rune
	path  string
	x     int
	y     int
	key   string
}

func coordinates(x int, y int, path string) (int, int) {
	re := regexp.MustCompile(`(se|sw|nw|ne|w|e)`)
	matches := re.FindAllStringSubmatch(path, -1)
	for _, group := range matches {
		switch group[0] {
		case "w":
			x -= 2
		case "sw":
			x--
			y += 2
		case "nw":
			x--
			y -= 2
		case "e":
			x += 2
		case "se":
			x++
			y += 2
		case "ne":
			x++
			y -= 2
		}
	}
	return x, y
}

func main() {
	path, _ := os.Getwd()
	file, err := os.Open(path + "/day24/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	pieces := []piece{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		x, y := coordinates(0, 0, line)
		p := piece{'b', line, x, y, fmt.Sprint(x, ",", y)}
		pieces = append(pieces, p)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	part2(part1(pieces))
}

func part1(pieces []piece) map[string]piece {
	occurrences := make(map[string]piece)
	for _, p := range pieces {
		current, exists := occurrences[p.key]
		if exists {
			switch current.color {
			case 'w':
				current.color = 'b'
			case 'b':
				current.color = 'w'
			}
			occurrences[p.key] = current
		} else {
			occurrences[p.key] = p
		}
	}
	count := 0
	for _, piece := range occurrences {
		if piece.color == 'b' {
			count++
		}
	}
	fmt.Println(count)
	return occurrences
}

func piecesAround(p piece, pieces map[string]piece) []piece {
	directions := []string{"e", "w", "nw", "sw", "ne", "se"}
	list := []piece{}
	for _, dir := range directions {
		x, y := coordinates(p.x, p.y, dir)
		other, exists := pieces[fmt.Sprint(x, ",", y)]
		if exists {
			list = append(list, other)
		} else {
			list = append(list, piece{'w', p.path + dir, x, y, fmt.Sprint(x, ",", y)})
		}
	}
	return list
}

func tryFlip(p piece, pieces map[string]piece) []piece {
	around := piecesAround(p, pieces)
	blacks := 0
	result := []piece{}
	for _, adj := range around {
		if adj.color == 'b' {
			blacks++
		}
		if p.color == 'b' && adj.color == 'w' {
			result = append(result, tryFlip(adj, pieces)...)
		}
	}
	switch {
	case p.color == 'w' && blacks == 2:
		p.color = 'b'
	case p.color == 'b' && (blacks == 0 || blacks > 2):
		p.color = 'w'
	}
	result = append(result, p)
	// fmt.Println(result)
	return result
}

func part2(pieces map[string]piece) {
	currentMap := pieces
	for i := 1; i <= 100; i++ {
		newMap := make(map[string]piece)
		for _, p := range currentMap {
			for _, piece := range tryFlip(p, currentMap) {
				newMap[piece.key] = piece
			}
		}
		currentMap = newMap
		count := 0
		for _, piece := range currentMap {
			if piece.color == 'b' {
				count++
			}
		}
		fmt.Println("day", i, " => ", count, "blacks")
	}
}
