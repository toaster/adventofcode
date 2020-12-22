package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
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
	done := false
	for !done {
		done = true
		for all, ings := range ingredientsByAllergen {
			if len(ings) == 1 {
				for all2 := range ingredientsByAllergen {
					if all2 == all {
						continue
					}
					ingredientsByAllergen[all2] = remove(ingredientsByAllergen[all2], ings[0])
				}
			} else {
				done = false
			}
		}
	}
	var allergens []string
	for all := range ingredientsByAllergen {
		allergens = append(allergens, all)
	}
	sort.Strings(allergens)
	var dangerousIngredients []string
	for _, allergen := range allergens {
		dangerousIngredients = append(dangerousIngredients, ingredientsByAllergen[allergen][0])
	}
	fmt.Println(strings.Join(dangerousIngredients, ","))
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

func remove(c []string, e string) []string {
	for i, s := range c {
		if s == e {
			return append(c[:i], c[i+1:]...)
		}
	}
	return c
}
