package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "failed reading standard input:", err)
		os.Exit(1)
	}

	ingredients := map[int][]string{}
	indexesByAllergen := map[string][]int{}
	for i, line := range strings.Split(strings.Trim(string(input), "\n"), "\n") {
		line = strings.Trim(line, ")")
		ingredientsAndAllergens := strings.Split(line, " (contains ")
		ingredients[i] = strings.Split(ingredientsAndAllergens[0], " ")
		for _, all := range strings.Split(ingredientsAndAllergens[1], ", ") {
			indexesByAllergen[all] = append(indexesByAllergen[all], i)
		}
	}
	ingredientsByAllergen := map[string][]string{}
	for all, indexes := range indexesByAllergen {
		for i, index := range indexes {
			if i == 0 {
				ingredientsByAllergen[all] = ingredients[index]
			} else {
				ingredientsByAllergen[all] = intersect(ingredientsByAllergen[all], ingredients[index])
			}
		}
	}
	count := 0
	for _, ings := range ingredients {
	Ing:
		for _, ing := range ings {
			for _, is := range ingredientsByAllergen {
				if contains(is, ing) {
					continue Ing
				}
			}
			count++
		}
	}
	fmt.Println("count:", count)
}

func add(coll []string, e string) []string {
	for _, s := range coll {
		if s == e {
			return coll
		}
	}
	return append(coll, e)
}

func contains(c []string, s string) bool {
	for _, e := range c {
		if e == s {
			return true
		}
	}
	return false
}

func intersect(a, b []string) []string {
	var c []string
	for _, s := range a {
		if contains(b, s) {
			c = append(c, s)
		}
	}
	return c
}
