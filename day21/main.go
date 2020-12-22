package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	path, _ := os.Getwd()
	file, err := os.Open(path + "/day21/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	ingredients := [][]string{}
	allergens := [][]string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = line[:len(line)-1]
		parts := strings.Split(line, " (contains ")
		ingredients = append(ingredients, strings.Split(parts[0], " "))
		allergens = append(allergens, strings.Split(parts[1], ", "))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	part1_2(ingredients, allergens)
}

func intersection(a, b []string) (c []string) {
	m := make(map[string]bool)

	for _, item := range a {
		m[item] = true
	}

	for _, item := range b {
		if _, ok := m[item]; ok {
			c = append(c, item)
		}
	}
	return
}

func hasSingleIngredient(candidateIngredients map[string][]string) bool {
	for _, ings := range candidateIngredients {
		if len(ings) > 1 {
			return false
		}
	}
	return true
}

func indexOf(list []string, item string) int {
	for idx, val := range list {
		if val == item {
			return idx
		}
	}
	return -1
}

func remove(list []string, index int) []string {
	if index >= 0 {
		copy(list[index:], list[index+1:])
		list[len(list)-1] = ""
		list = list[:len(list)-1]
	}
	return list
}

func part1_2(ingredients [][]string, allergens [][]string) {
	candidateIngredients := make(map[string][]string)
	for i := 0; i < len(ingredients); i++ {
		for _, allergen := range allergens[i] {
			_, exists := candidateIngredients[allergen]
			if !exists {
				candidateIngredients[allergen] = ingredients[i]
			} else {
				candidateIngredients[allergen] = intersection(candidateIngredients[allergen], ingredients[i])
			}
		}
	}
	for !hasSingleIngredient(candidateIngredients) {
		for allergen, ings := range candidateIngredients {
			if len(ings) == 1 {
				for otherAllergen, otherIngs := range candidateIngredients {
					if otherAllergen != allergen {
						index := indexOf(otherIngs, ings[0])
						candidateIngredients[otherAllergen] = remove(otherIngs, index)
					}
				}
			}
		}
	}
	allIngredients := []string{}
	for _, ings := range ingredients {
		for _, ing := range ings {
			allIngredients = append(allIngredients, ing)
		}
	}
	allAllergens := []string{}
	for allergen, ings := range candidateIngredients {
		allAllergens = append(allAllergens, allergen)
		index := indexOf(allIngredients, ings[0])
		for index >= 0 {
			allIngredients = remove(allIngredients, index)
			index = indexOf(allIngredients, ings[0])
		}
	}
	// part1
	fmt.Println(len(allIngredients))

	// part2
	sort.Strings(allAllergens)
	dangerousIngs := []string{}
	for _, allergen := range allAllergens {
		dangerousIngs = append(dangerousIngs, candidateIngredients[allergen][0])
	}
	fmt.Println(strings.Join(dangerousIngs, ","))

}
